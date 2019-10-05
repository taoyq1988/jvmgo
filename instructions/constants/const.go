package constants

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

// xconst: push x
type Const struct {
	base.NoOperandsInstruction
	K heap.Slot
	L bool
}

func (constant *Const) Execute(frame *rtda.Frame) {
	frame.Push(constant.K)
	if constant.L {
		frame.PushNull()
	}
}
