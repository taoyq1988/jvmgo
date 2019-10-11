package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IMul struct {
	base.NoOperandsInstruction
}

func (_ *IMul) Execute(frame *rtda.Frame) {
	val1 := frame.PopInt()
	val2 := frame.PopInt()
	frame.PushInt(val1 * val2)
}

type LMul struct {
	base.NoOperandsInstruction
}

func (_ *LMul) Execute(frame *rtda.Frame) {
	val1 := frame.PopLong()
	val2 := frame.PopLong()
	frame.PushLong(val1 * val2)
}

type FMul struct {
	base.NoOperandsInstruction
}

func (_ *FMul) Execute(frame *rtda.Frame) {
	val1 := frame.PopFloat()
	val2 := frame.PopFloat()
	frame.PushFloat(val1 * val2)
}

type DMul struct {
	base.NoOperandsInstruction
}

func (_ *DMul) Execute(frame *rtda.Frame) {
	val1 := frame.PopDouble()
	val2 := frame.PopDouble()
	frame.PushDouble(val1 * val2)
}
