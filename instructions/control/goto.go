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

// Branch always (wide index)
type GotoW struct {
	offset int
}

func (instr *GotoW) FetchOperands(reader *base.CodeReader) {
	instr.offset = int(reader.ReadInt32())
}
func (instr *GotoW) Execute(frame *rtda.Frame) {
	base.Branch(frame, instr.offset)
}
