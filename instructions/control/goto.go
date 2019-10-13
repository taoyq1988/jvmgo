package control

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type Goto struct {
	base.BranchInstruction
}

func (gt *Goto) Execute(frame *rtda.Frame) {
	base.Branch(frame, gt.Offset)
}
