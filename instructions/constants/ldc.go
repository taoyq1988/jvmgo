package constants

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type LDC struct {
	base.Index8Instruction
}

func (ldc *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, ldc.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (ldc *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, ldc.Index)
}

func _ldc(frame *rtda.Frame, index uint) {
	c := frame.GetConstantPool().GetConstant(index)
	switch x := c.(type) {
	case int32:
		frame.PushInt(x)
	case float32:
		frame.PushFloat(x)
	case *heap.ConstantString:
		frame.PushRef(x.GetJString())
	case *heap.ConstantClass:
		frame.PushRef(x.GetClass().JClass)
	default:
		//todo ref to MethodType and MethodHandle
		panic("ldc!")
	}
}

type LDC2_W struct {
	base.Index16Instruction
}

func (ldc2 *LDC2_W) Execute(frame *rtda.Frame) {
	c := frame.GetConstantPool().GetConstant(ldc2.Index)
	switch x := c.(type) {
	case int64:
		frame.PushLong(x)
	case float64:
		frame.PushDouble(x)
	default:
		panic("ldc2_w!")
	}
}
