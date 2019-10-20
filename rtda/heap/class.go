package heap

import (
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/classpath"
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
}

func (class *Class) String() string {
	return `{Class name:` + class.Name + `}`
}

/**
Create class
 */
func newClass(cf *classfile.Classfile) *Class {
	class := &Class{}
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
