package heap

import (
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/classpath"
	"sync"
)

// initialization state
const (
	_notInitialized   = 0 // This Class object is verified and prepared but not initialized.
	_beingInitialized = 1 // This Class object is being initialized by some particular thread T.
	_fullyInitialized = 2 // This Class object is fully initialized and ready for use.
	_initFailed       = 3 // This Class object is in an erroneous state, perhaps because initialization was attempted and failed.
)

type ClassMember struct {
	AccessFlag
	Name           string
	Descriptor     string
	Signature      string
	AnnotationData []byte // RuntimeVisibleAnnotations_attribute
	Class          *Class
}

func (member *ClassMember) parseMemberData(cf *classfile.Classfile, cfMember classfile.MemberInfo) {
	member.AccessFlag = AccessFlag(cfMember.AccessFlags)
	member.Name = cf.GetUTF8(cfMember.NameIndex)
	member.Descriptor = cf.GetUTF8(cfMember.DescriptorIndex)
	member.Signature = cf.GetUTF8(cfMember.GetSignatureIndex())
	member.AnnotationData = cfMember.GetRuntimeVisibleAnnotationsAttributeData()
}

type ClassAttributes struct {
	SourceFile      string
	Signature       string
	AnnotationData  []byte // RuntimeVisibleAnnotations_attribute
	EnclosingMethod *EnclosingMethod
}

type EnclosingMethod struct {
	ClassName        string
	MethodName       string
	MethodDescriptor string
}

type Class struct {
	AccessFlag
	ClassAttributes
	ConstantPool
	Name               string // thisClassName
	superClassName     string
	interfaceNames     []string
	Fields             []*Field
	Methods            []*Method
	instanceFieldCount uint
	staticFieldCount   uint
	StaticFieldSlots   []Slot
	vtable             []*Method // virtual method table
	JClass             *Object   // java.lang.Class instance
	SuperClass         *Class
	Interfaces         []*Class
	LoadedFrom         classpath.Entry
	initState          int
	InitCond           *sync.Cond
	initThread         uintptr
}

func (class *Class) String() string {
	return `{Class name:` + class.Name + `}`
}

func (class *Class) NameJlsFormat() string {
	return SlashToDot(class.Name)
}

func (class *Class) NewObj() *Object {
	if class.instanceFieldCount > 0 {
		fields := make([]Slot, class.instanceFieldCount)
		obj := newObj(class, fields, nil)
		obj.initFields()
		return obj
	}
	return newObj(class, nil, nil)
}

func (class *Class) InitializationNotStarted() bool {
	return class.initState < _beingInitialized // todo
}
func (class *Class) IsBeingInitialized() (bool, uintptr) {
	return class.initState == _beingInitialized, class.initThread
}
func (class *Class) IsFullyInitialized() bool {
	return class.initState == _fullyInitialized
}
func (class *Class) IsInitializationFailed() bool {
	return class.initState == _initFailed
}
func (class *Class) MarkBeingInitialized(thread uintptr) {
	class.initState = _beingInitialized
	class.initThread = thread
}
func (class *Class) MarkFullyInitialized() {
	class.initState = _fullyInitialized
}

func (class *Class) getField(name, descriptor string, isStatic bool) *Field {
	for k := class; k != nil; k = k.SuperClass {
		for _, field := range k.Fields {
			if field.IsStatic() == isStatic && field.Name == name && field.Descriptor == descriptor {
				return field
			}
		}
	}
	return nil
}

/**
GetMethod
 */
func (class *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for k := class; k != nil; k = k.SuperClass {
		for _, method := range k.Methods {
			if method.IsStatic() == isStatic &&
				method.Name == name &&
				method.Descriptor == descriptor {
				return method
			}
		}
	}
	// todo
	return nil
}

func (class *Class) getDeclaredMethod(name, descriptor string, isStatic bool) *Method {
	for _, method := range class.Methods {
		if method.IsStatic() == isStatic && method.Name == name && method.Descriptor == descriptor {
			return method
		}
	}
	return nil
}

func (class *Class) GetStaticMethod(name, descriptor string) *Method {
	return class.getMethod(name, descriptor, true)
}

func (class *Class) GetClinitMethod() *Method {
	return class.getDeclaredMethod(clinitMethodName, clinitMethodDesc, true)
}


/**
judge function
 */
func (class *Class) IsArray() bool {
	return class.Name[0] == '['
}

func (class *Class) isJlObject() bool {
	return class == _jlObjectClass
}
func (class *Class) isJlCloneable() bool {
	return class == _jlCloneableClass
}
func (class *Class) isJioSerializable() bool {
	return class == _ioSerializableClass
}

func (class *Class) ComponentClass() *Class {
	componentClassName := getComponentClassName(class.Name)
	return bootLoader.LoadClass(componentClassName)
}


/**
isinstanceof && cast
*/
func (class *Class) IsAssignableFrom(cls *Class) bool {
	if class == cls && class.isSuperClassOf(cls) && class.isSuperInterfaceOf(cls) {
		return true
	}
	return false
}

func (class *Class) IsImplements(iface *Class) bool {
	for k := class; k != nil; k = k.SuperClass {
		for _, i := range k.Interfaces {
			if i == iface || i.isSubInterfaceOf(iface) {
				return true
			}
		}
	}
	return false
}

func (class *Class) isSuperInterfaceOf(iface *Class) bool {
	return iface.isSubInterfaceOf(class)
}

func (class *Class) isSubInterfaceOf(iface *Class) bool {
	for _, superInterface := range class.Interfaces {
		if superInterface == iface || superInterface.isSubInterfaceOf(iface) {
			return true
		}
	}
	return false
}

func (class *Class) isSuperClassOf(c *Class) bool {
	return c.isSubClassOf(class)
}

func (class *Class) isSubClassOf(c *Class) bool {
	for k := class.SuperClass; k != nil; k = k.SuperClass {
		if k == c {
			return true
		}
	}
	return false
}

/**
Create class
*/
func newClass(cf *classfile.Classfile) *Class {
	class := &Class{}
	class.InitCond = sync.NewCond(&sync.Mutex{})
	class.AccessFlag = AccessFlag(cf.AccessFlags)
	class.parseConstantPool(cf)
	class.parseClassNames(cf)
	class.parseFields(cf)
	class.parseMethods(cf)
	class.parseAttributes(cf)
	return class
}

func (class *Class) parseConstantPool(cf *classfile.Classfile) {
	class.ConstantPool = newConstantPool(cf)
}

func (class *Class) parseClassNames(cf *classfile.Classfile) {
	class.Name = cf.GetThisClassName()
	class.superClassName = cf.GetSuperClassName()
	class.interfaceNames = cf.GetInterfaceNames()
}

func (class *Class) parseFields(cf *classfile.Classfile) {
	class.Fields = make([]*Field, len(cf.Fields))
	for i, fieldInfo := range cf.Fields {
		class.Fields[i] = newField(class, cf, fieldInfo)
	}
}

func (class *Class) parseMethods(cf *classfile.Classfile) {
	class.Methods = make([]*Method, len(cf.Methods))
	for i, methodInfo := range cf.Methods {
		class.Methods[i] = newMethod(class, cf, methodInfo)
		class.Methods[i].Slot = uint(i)
	}
}

func (class *Class) parseAttributes(cf *classfile.Classfile) {
	class.SourceFile = cf.GetUTF8(cf.GetSourceFileIndex())
	class.Signature = cf.GetUTF8(cf.GetSignatureIndex())
	class.AnnotationData = cf.GetRuntimeVisibleAnnotationsAttributeData()
	class.EnclosingMethod = getEnclosingMethod(cf)
}

func getEnclosingMethod(cf *classfile.Classfile) *EnclosingMethod {
	if emAttr, found := cf.GetEnclosingMethodAttribute(); found {
		methodName, methodDescriptor := getNameAndType(cf, emAttr.MethodIndex)
		return &EnclosingMethod{
			ClassName:        cf.GetClassNameOf(emAttr.ClassIndex),
			MethodName:       methodName,
			MethodDescriptor: methodDescriptor,
		}
	}
	return nil
}
