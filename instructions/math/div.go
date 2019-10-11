package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IDiv struct {
	base.NoOperandsInstruction
}

func (_ *IDiv) Execute(frame *rtda.Frame) {
	val1 := frame.PopInt()
	val2 := frame.PopInt()
	if val1 == 0 {
		//todo
	}
	frame.PushInt(val2 / val1)
}

type LDiv struct {
	base.NoOperandsInstruction
}

func (_ *LDiv) Execute(frame *rtda.Frame) {
	val1 := frame.PopLong()
	val2 := frame.PopLong()
	if val1 == 0 {
		//todo
	}
	frame.PushLong(val2 / val1)
}

type FDiv struct {
	base.NoOperandsInstruction
}

func (_ *FDiv) Execute(frame *rtda.Frame) {
	val1 := frame.PopFloat()
	val2 := frame.PopFloat()
	if val1 == 0 {
		//todo
	}
	frame.PushFloat(val2 / val1)
}

type DDiv struct {
	base.NoOperandsInstruction
}

func (_ *DDiv) Execute(frame *rtda.Frame) {
	val1 := frame.PopDouble()
	val2 := frame.PopDouble()
	if val1 == 0 {
		//todo
	}
	frame.PushDouble(val2 / val1)
}
