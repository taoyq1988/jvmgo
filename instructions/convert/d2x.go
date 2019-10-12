package convert

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type D2I struct {
	base.NoOperandsInstruction
}

func (_ *D2I) Execute(frame *rtda.Frame) {
	d := frame.PopDouble()
	frame.PushInt(int32(d))
}

type D2L struct {
	base.NoOperandsInstruction
}

func (_ *D2L) Execute(frame *rtda.Frame) {
	d := frame.PopDouble()
	frame.PushLong(int64(d))
}

type D2F struct {
	base.NoOperandsInstruction
}

func (_ *D2F) Execute(frame *rtda.Frame) {
	d := frame.PopDouble()
	frame.PushFloat(float32(d))
}
