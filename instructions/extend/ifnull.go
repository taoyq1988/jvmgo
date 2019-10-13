package extended

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

// Branch if reference is null
type IfNull struct{ base.BranchInstruction }

func (instr *IfNull) Execute(frame *rtda.Frame) {
	ref := frame.PopRef()
	if ref == nil {
		base.Branch(frame, instr.Offset)
	}
}

// Branch if reference not null
type IfNonNull struct{ base.BranchInstruction }

func (instr *IfNonNull) Execute(frame *rtda.Frame) {
	ref := frame.PopRef()
	if ref != nil {
		base.Branch(frame, instr.Offset)
	}
}
