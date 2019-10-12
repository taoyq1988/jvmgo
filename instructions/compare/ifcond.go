package compare

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IfEQ struct {
	base.BranchInstruction
}

func (ifc *IfEQ) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val == 0 {
		base.Branch(frame, ifc.Offset)
	}
}

type IfNE struct {
	base.BranchInstruction
}

func (ifc *IfNE) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val != 0 {
		base.Branch(frame, ifc.Offset)
	}
}

type IfLT struct {
	base.BranchInstruction
}

func (ifc *IfLT) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val < 0 {
		base.Branch(frame, ifc.Offset)
	}
}

type IfGE struct {
	base.BranchInstruction
}

func (ifc *IfGE) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val >= 0 {
		base.Branch(frame, ifc.Offset)
	}
}

type IfGT struct {
	base.BranchInstruction
}

func (ifc *IfGT) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val > 0 {
		base.Branch(frame, ifc.Offset)
	}
}

type IfLE struct {
	base.BranchInstruction
}

func (ifc *IfLE) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	if val <= 0 {
		base.Branch(frame, ifc.Offset)
	}
}
