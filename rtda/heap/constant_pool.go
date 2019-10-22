package heap

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/classfile"
)

type Constant interface{}
type ConstantPool []Constant

func newConstantPool(cf *classfile.Classfile) ConstantPool {
	cfCp := cf.ConstPool
	rtCp := make([]Constant, len(cfCp))
	for i := 1; i < len(cfCp); i++ {
		cpInfo := cfCp[i]
		switch x := cpInfo.(type) {
		case string:
			rtCp[i] = cpInfo
		case int32, float32:
			rtCp[i] = cpInfo
		case int64, float64:
			rtCp[i] = cpInfo
			i++
		case classfile.ConstantStringInfo:
			rtCp[i] = newConstantString(cf.GetUTF8(x.StringIndex))
		case classfile.ConstantClassInfo:
			rtCp[i] = newConstantClass(cf, x)
		case classfile.ConstantFieldRefInfo:
			rtCp[i] = newConstantFieldRef(cf, x)
		case classfile.ConstantMethodRefInfo:
			rtCp[i] = newConstantMethodRef(cf, x)
		case classfile.ConstantInterfaceMethodRefInfo:
			rtCp[i] = newConstantInterfaceMethodRef(cf, x)
		case classfile.ConstantInvokeDynamicInfo:
			rtCp[i] = newConstantInvokeDynamic(cf, rtCp, x)
		case classfile.ConstantMethodHandleInfo:
			rtCp[i] = newConstantMethodHandle(x)
		case classfile.ConstantMethodTypeInfo:
			rtCp[i] = newConstantMethodType(x)
		default:
			// todo
			// panic
		}
	}

	return rtCp
}

func (cp ConstantPool) GetConstantString(index uint) *ConstantString {
	return cp.GetConstant(index).(*ConstantString)
}

func (cp ConstantPool) GetConstantClass(index uint) *ConstantClass {
	return cp.GetConstant(index).(*ConstantClass)
}

func (cp ConstantPool) GetConstantFieldRef(index uint) *ConstantFieldRef {
	return cp.GetConstant(index).(*ConstantFieldRef)
}

func (cp ConstantPool) GetConstant(index uint) Constant {
	// TODO: check index
	return cp[index]
}

/**
ConstantString
*/
type ConstantString struct {
	goStr string
	jStr  *Object
}

func newConstantString(str string) *ConstantString {
	return &ConstantString{goStr: str}
}

func (s *ConstantString) GetJString() *Object {
	if s.jStr == nil {
		s.jStr = JSFromGoStr(s.goStr)
	}
	return s.jStr
}

/**
ConstantClass
*/
type ConstantClass struct {
	name  string
	class *Class
}

func newConstantClass(cf *classfile.Classfile, cfc classfile.ConstantClassInfo) *ConstantClass {
	return &ConstantClass{
		name: cf.GetUTF8(cfc.NameIndex),
	}
}

func (cr *ConstantClass) GetClass() *Class {
	if cr.class == nil {
		cr.resolve()
	}
	return cr.class
}

// todo
func (cr *ConstantClass) resolve() {
	// load class
	cr.class = bootLoader.LoadClass(cr.name)
}

/**
ConstantMemberRef
*/
type ConstantMemberRef struct {
	className  string
	name       string
	descriptor string
}

func (mr *ConstantMemberRef) init(cf *classfile.Classfile, classIdx, nameAndTypeIdx uint16) {
	mr.className = cf.GetClassNameOf(classIdx)
	mr.name, mr.descriptor = getNameAndType(cf, nameAndTypeIdx)
}

/**
ConstantFieldRef
*/
type ConstantFieldRef struct {
	ConstantMemberRef
	field *Field
}

func newConstantFieldRef(cf *classfile.Classfile, cfRef classfile.ConstantFieldRefInfo) *ConstantFieldRef {

	ref := &ConstantFieldRef{}
	ref.init(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	return ref
}

func (fr *ConstantFieldRef) String() string {
	return fmt.Sprintf("{ConstantFieldref className:%v name:%v descriptor:%v}",
		fr.className, fr.name, fr.descriptor)
}

func (fr *ConstantFieldRef) GetField(static bool) *Field {
	if fr.field == nil {
		if static {
			fr.resolveStaticField()
		} else {
			fr.resolveInstanceField()
		}
	}
	return fr.field
}

func (fr *ConstantFieldRef) resolveInstanceField() {
	//fromClass := bootLoader.LoadClass(fr.className)
	//
	//for class := fromClass; class != nil; class = class.SuperClass {
	//	field := class.getField(fr.name, fr.descriptor, false)
	//	if field != nil {
	//		fr.field = field
	//		return
	//	}
	//}

	// todo
	panic(fmt.Errorf("instance field not found! %v", fr))
}

func (fr *ConstantFieldRef) resolveStaticField() {
	//fromClass := bootLoader.LoadClass(fr.className)
	//
	//for class := fromClass; class != nil; class = class.SuperClass {
	//	field := class.getField(fr.name, fr.descriptor, true)
	//	if field != nil {
	//		fr.field = field
	//		return
	//	}
	//	if fr._findInterfaceField(class) {
	//		return
	//	}
	//}

	// todo
	panic(fmt.Errorf("static field not found! %v", fr))
}

func (fr *ConstantFieldRef) _findInterfaceField(class *Class) bool {
	for _, iface := range class.Interfaces {
		for _, f := range iface.Fields {
			if f.Name == fr.name && f.Descriptor == fr.descriptor {
				fr.field = f
				return true
			}
		}
		if fr._findInterfaceField(iface) {
			return true
		}
	}
	return false
}

/**
ConstantMethodRef
*/
type ConstantMethodRef struct {
	ConstantMemberRef
	ParamSlotCount uint
	method         *Method
	vslot          int
}

func newConstantMethodRef(cf *classfile.Classfile, cfRef classfile.ConstantMethodRefInfo) *ConstantMethodRef {

	ref := &ConstantMethodRef{vslot: -1}
	ref.init(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	ref.ParamSlotCount = calcParamSlotCount(ref.descriptor)
	return ref
}

func (mr *ConstantMethodRef) GetMethod(static bool) *Method {
	if mr.method == nil {
		if static {
			mr.resolveStaticMethod()
		} else {
			mr.resolveSpecialMethod()
		}
	}
	return mr.method
}

func (mr *ConstantMethodRef) resolveStaticMethod() {
	method := mr.findMethod(true)
	if method != nil {
		mr.method = method
	} else {
		// todo
		panic("static method not found!")
	}
}

func (mr *ConstantMethodRef) resolveSpecialMethod() {
	method := mr.findMethod(false)
	if method != nil {
		mr.method = method
		return
	}

	// todo
	// class := mr.cp.class.classLoader.LoadClass(mr.className)
	// if class.IsInterface() {
	// 	method = mr.findMethodInInterfaces(class)
	// 	if method != nil {
	// 		mr.method = method
	// 		return
	// 	}
	// }

	// todo
	panic("special method not found!")
}

func (mr *ConstantMethodRef) findMethod(isStatic bool) *Method {
	//class := bootLoader.LoadClass(mr.className)
	//return class.getMethod(mr.name, mr.descriptor, isStatic)
	return nil
}

// todo
/*func (mr *ConstantMethodref) findMethodInInterfaces(iface *Class) *Method {
	for _, m := range iface.methods {
		if !m.IsAbstract() {
			if m.name == mr.name && m.descriptor == mr.descriptor {
				return m
			}
		}
	}

	for _, superIface := range iface.interfaces {
		if m := mr.findMethodInInterfaces(superIface); m != nil {
			return m
		}
	}

	return nil
}*/

func (mr *ConstantMethodRef) GetVirtualMethod(ref *Object) *Method {
	//if mr.vslot < 0 {
	//	mr.vslot = getVslot(ref.Class, mr.name, mr.descriptor)
	//}
	//return ref.Class.vtable[mr.vslot]
	return nil
}

/**
ConstantInterfaceMethodRef
*/
type ConstantInterfaceMethodRef struct {
	ConstantMethodRef
}

func newConstantInterfaceMethodRef(cf *classfile.Classfile, cfRef classfile.ConstantInterfaceMethodRefInfo) *ConstantInterfaceMethodRef {
	ref := &ConstantInterfaceMethodRef{}
	ref.init(cf, cfRef.ClassIndex, cfRef.NameAndTypeIndex)
	ref.ParamSlotCount = calcParamSlotCount(ref.descriptor)
	return ref
}

// todo
func (imr *ConstantInterfaceMethodRef) FindInterfaceMethod(ref *Object) *Method {
	for class := ref.Class; class != nil; class = class.SuperClass {
		method := class.getMethod(imr.name, imr.descriptor, false)
		if method != nil {
			return method
		}
	}

	if method := findInterfaceMethod(ref.Class.Interfaces, imr.name, imr.descriptor); method != nil {
		return method
	} else {
		//TODO
		panic("virtual method not found!")
	}
}

func findInterfaceMethod(interfaces []*Class, name, descriptor string) *Method {
	for i := 0; i < len(interfaces); i++ {
		if method := findInterfaceMethod(interfaces[i].Interfaces, name, descriptor); method != nil {
			return method
		}
		method := interfaces[i].getMethod(name, descriptor, false)
		if method != nil {
			return method
		}
	}
	return nil
}

/**
ConstantInvokeDynamic
*/
type ConstantInvokeDynamic struct {
	name               string
	_type              string
	bootstrapMethodRef uint16 // method handle
	bootstrapArguments []uint16
	cp                 ConstantPool
}

func newConstantInvokeDynamic(cf *classfile.Classfile, cp ConstantPool, indyInfo classfile.ConstantInvokeDynamicInfo) *ConstantInvokeDynamic {
	name, _type := getNameAndType(cf, indyInfo.NameAndTypeIndex)
	bm := cf.GetBootstrapMethods()[indyInfo.BootstrapMethodAttrIndex]
	return &ConstantInvokeDynamic{
		name:               name,
		_type:              _type,
		bootstrapMethodRef: bm.BootstrapMethodRef,
		bootstrapArguments: bm.BootstrapArguments,
		cp:                 cp,
	}
}

func (indy *ConstantInvokeDynamic) MethodHandle() {
	kMH := indy.cp.GetConstant(uint(indy.bootstrapMethodRef)).(*ConstantMethodHandle)
	println(kMH)
}

/**
ConstantMethodHandle
*/
type ConstantMethodHandle struct {
	referenceKind  uint8
	referenceIndex uint16
}

func newConstantMethodHandle(mhInfo classfile.ConstantMethodHandleInfo) *ConstantMethodHandle {
	return &ConstantMethodHandle{
		referenceKind:  mhInfo.ReferenceKind,
		referenceIndex: mhInfo.ReferenceIndex,
	}
}

/**
ConstantMethodType
*/
type ConstantMethodType struct {
	// todo
}

func newConstantMethodType(mtInfo classfile.ConstantMethodTypeInfo) *ConstantMethodType {
	return &ConstantMethodType{
		// todo
	}
}
