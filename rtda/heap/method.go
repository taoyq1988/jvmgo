package heap

import "github.com/taoyq1988/jvmgo/classfile"

const (
	mainMethodName   = "main"
	mainMethodDesc   = "([Ljava/lang/String;)V"
	clinitMethodName = "<clinit>"
	clinitMethodDesc = "()V"
	constructorName  = "<init>"
)

type MethodData struct {
	MaxStack                uint
	MaxLocals               uint
	Code                    []byte
	exceptionTable          []classfile.ExceptionTableEntry
	lineNumberTable         []classfile.LineNumberTableEntry
	ParameterAnnotationData []byte // RuntimeVisibleParameterAnnotations_attribute
	AnnotationDefaultData   []byte // AnnotationDefault_attribute
}

type Method struct {
	ClassMember
	MethodData
	ParsedDescriptor
	ParamCount          uint
	ParamSlotCount      uint
	Slot                uint
	exceptionTableIndex []uint16
	nativeMethod        interface{} // cannot use package 'native' because of cycle import!
	Instructions        interface{} // []instructions.Instruction
}

func newMethod(class *Class, cf *classfile.Classfile, cfMember classfile.MemberInfo) *Method {
	method := &Method{}
	method.Class = class
	method.parseMemberData(cf, cfMember)
	method.parseAttributes(cf, cfMember)
	method.parseDescriptor()
	return method
}

func (method *Method) parseAttributes(cf *classfile.Classfile, cfMember classfile.MemberInfo) {
	if codeAttr, f := cfMember.GetCodeAttribute(); f {
		method.exceptionTableIndex = cfMember.GetExceptionIndexTable()
		method.MaxStack = uint(codeAttr.MaxStack)
		method.MaxLocals = uint(codeAttr.MaxLocals)
		method.Code = codeAttr.Code
		method.exceptionTable = codeAttr.ExceptionTable
		method.lineNumberTable = codeAttr.GetLineNumberTable()
	}
	method.ParameterAnnotationData = cfMember.GetRuntimeVisibleParameterAnnotationsAttributeData()
	method.AnnotationDefaultData = cfMember.GetAnnotationDefaultAttributeData()
}

func (method *Method) parseDescriptor() {
	method.ParsedDescriptor = parseMethodDescriptor(method.Descriptor)
	method.ParamCount = uint(len(method.ParameterTypes))
	method.ParamSlotCount = method.getParamSlotCount()
	if !method.IsStatic() {
		method.ParamSlotCount++
	}
}

/**
class vtable
*/
func getVSlot(class *Class, name, descriptor string) int {
	for i, m := range class.vtable {
		if m.Name == name && m.Descriptor == descriptor {
			return i
		}
	}
	return -1
}

func createVTable(class *Class) {
	class.vtable = copySuperVTable(class)
	for _, m := range class.Methods {
		if isVirtualMethod(m) {
			if i := indexOf(class.vtable, m); i != -1 {
				class.vtable[i] = m //override
			} else {
				addVMethod(class, m)
			}
		}
	}
	eachInterfaceMethod(class, func(method *Method) {
		if i := indexOf(class.vtable, method); i == -1 {
			addVMethod(class, method)
		}
	})
}

func copySuperVTable(class *Class) []*Method {
	if class.SuperClass != nil {
		superVTable := class.SuperClass.vtable
		newVTable := make([]*Method, len(superVTable))
		copy(newVTable, superVTable)
		return newVTable
	}
	return nil
}

func isVirtualMethod(method *Method) bool {
	return !method.IsStatic() &&
		!method.IsPrivate() &&
		method.Name != constructorName
}

func indexOf(vTable []*Method, method *Method) int {
	for i, m := range vTable {
		if m.Name == method.Name && m.Descriptor == method.Descriptor {
			return i
		}
	}
	return -1
}

func addVMethod(class *Class, method *Method) {
	_len := len(class.vtable)
	if _len == cap(class.vtable) {
		newVTable := make([]*Method, _len, _len+8)
		copy(newVTable, class.vtable)
		class.vtable = newVTable
	}
	class.vtable = append(class.vtable, method)
}

func eachInterfaceMethod(class *Class, f func(*Method)) {
	for _, iface := range class.Interfaces {
		eachInterfaceMethod(iface, f)
		for _, m := range iface.Methods {
			f(m)
		}
	}
}
