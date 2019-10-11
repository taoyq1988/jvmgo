package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type INeg struct {
	base.NoOperandsInstruction
}

func (_ INeg) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	frame.PushInt(-val)
}

type LNeg struct {
	base.NoOperandsInstruction
}

func (_ LNeg) Execute(frame *rtda.Frame) {
	val := frame.PopLong()
	frame.PushLong(-val)
}

type FNeg struct {
	base.NoOperandsInstruction
}

func (_ FNeg) Execute(frame *rtda.Frame) {
	val := frame.PopFloat()
	frame.PushFloat(-val)
}

type DNeg struct {
	base.NoOperandsInstruction
}

func (_ DNeg) Execute(frame *rtda.Frame) {
	val := frame.PopDouble()
	frame.PushDouble(-val)
}
