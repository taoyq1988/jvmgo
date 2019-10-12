package convert

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type F2I struct {
	base.NoOperandsInstruction
}

func (_ *F2I) Execute(frame *rtda.Frame) {
	f := frame.PopFloat()
	frame.PushInt(int32(f))
}

type F2L struct {
	base.NoOperandsInstruction
}

func (_ *F2L) Execute(frame *rtda.Frame) {
	f := frame.PopFloat()
	frame.PushLong(int64(f))
}

type F2D struct {
	base.NoOperandsInstruction
}

func (_ *F2D) Execute(frame *rtda.Frame) {
	f := frame.PopFloat()
	frame.PushDouble(float64(f))
}
