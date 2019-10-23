package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type InstanceOf struct {
	base.Index16Instruction
}

func (inst *InstanceOf) Execute(frame *rtda.Frame) {
	ref := frame.PopRef()
	cp := frame.GetConstantPool()
	kClass := cp.GetConstantClass(inst.Index)
	class := kClass.GetClass()

	if ref == nil {
		frame.PushInt(0)
	} else if ref.IsInstanceOf(class) {
		frame.PushInt(1)
	} else {
		frame.PushInt(0)
	}
}

type CheckCast struct {
	base.Index16Instruction
	class *heap.Class
}

func (cast *CheckCast) Execute(frame *rtda.Frame) {
	if cast.class == nil {
		cp := frame.GetConstantPool()
		kClass := cp.GetConstantClass(cast.Index)
		cast.class = kClass.GetClass()
	}
	ref := frame.PopRef()
	frame.PushRef(ref)
	if ref == nil {
		return
	}
	if !ref.IsInstanceOf(cast.class) {
		frame.Thread.ThrowClassCastException(ref.Class, cast.class)
	}
}
