package convert

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type L2I struct {
	base.NoOperandsInstruction
}

func (_ *L2I) Execute(frame *rtda.Frame) {
	l := frame.PopLong()
	frame.PushInt(int32(l))
}

type L2F struct {
	base.NoOperandsInstruction
}

func (_ *L2F) Execute(frame *rtda.Frame) {
	l := frame.PopLong()
	frame.PushFloat(float32(l))
}

type L2D struct {
	base.NoOperandsInstruction
}

func (_ *L2D) Execute(frame *rtda.Frame) {
	l := frame.PopLong()
	frame.PushDouble(float64(l))
}
