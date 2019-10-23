package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type PutField struct {
	base.Index16Instruction
	field *heap.Field
}

func (putF *PutField) Execute(frame *rtda.Frame) {
	if putF.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(putF.Index)
		putF.field = fieldRef.GetField(false)
	}

	val := frame.PopL(putF.field.IsLongOrDouble)
	ref := frame.PopRef()
	if ref == nil {
		frame.Thread.ThrowNEP()
		return
	}
	putF.field.PutValue(ref, val)
}

type GetField struct {
	base.Index16Instruction
	field *heap.Field
}

func (getF *GetField) Execute(frame *rtda.Frame) {
	if getF.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(getF.Index)
		getF.field = fieldRef.GetField(false)
	}

	ref := frame.PopRef()
	if ref == nil {
		frame.Thread.ThrowNEP()
		return
	}
	val := getF.field.GetValue(ref)
	frame.PushL(val, getF.field.IsLongOrDouble)
}
