package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IAdd struct {
	base.NoOperandsInstruction
}

func (_ *IAdd) Execute(frame *rtda.Frame) {
	v1 := frame.PopInt()
	v2 := frame.PopInt()
	frame.PushInt(v1 + v2)
}

type LAdd struct {
	base.NoOperandsInstruction
}

func (_ *LAdd) Execute(frame *rtda.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	frame.PushLong(v1 + v2)
}

type FAdd struct {
	base.NoOperandsInstruction
}

func (_ *FAdd) Execute(frame *rtda.Frame) {
	v1 := frame.PopFloat()
	v2 := frame.PopFloat()
	frame.PushFloat(v1 + v2)
}

type DAdd struct {
	base.NoOperandsInstruction
}

func (_ *DAdd) Execute(frame *rtda.Frame) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
	frame.PushDouble(v1 + v2)
}
