package stack

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type Pop struct {
	base.NoOperandsInstruction
}

func (pop *Pop) Execute(frame *rtda.Frame) {
	frame.Pop()
}

type Pop2 struct {
	base.NoOperandsInstruction
}

func (pop2 *Pop2) Execute(frame *rtda.Frame) {
	frame.Pop()
	frame.Pop()
}
