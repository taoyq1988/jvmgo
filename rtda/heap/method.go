package heap

import "github.com/taoyq1988/jvmgo/classfile"

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
	ParamCount     uint
	ParamSlotCount uint
	Slot           uint
	exceptionTableIndex []uint16
	nativeMethod   interface{} // cannot use package 'native' because of cycle import!
	Instructions   interface{} // []instructions.Instruction
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
