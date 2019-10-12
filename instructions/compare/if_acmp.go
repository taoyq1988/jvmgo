package compare

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IfACmpEQ struct {
	base.BranchInstruction
}

func (ifac *IfACmpEQ) Execute(frame *rtda.Frame) {
	if acmp(frame) {
		base.Branch(frame, ifac.Offset)
	}
}

type IfACmpNE struct {
	base.BranchInstruction
}

func (ifac *IfACmpNE) Execute(frame *rtda.Frame) {
	if !acmp(frame) {
		base.Branch(frame, ifac.Offset)
	}
}

func acmp(frame *rtda.Frame) bool {
	return frame.PopRef() == frame.PopRef()
}
