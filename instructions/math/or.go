package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IOr struct {
	base.NoOperandsInstruction
}

func (_ IOr) Execute(frame *rtda.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	frame.PushInt(v1 | v2)
}

type LOr struct {
	base.NoOperandsInstruction
}

func (_ LOr) Execute(frame *rtda.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	frame.PushLong(v1 | v2)
}
