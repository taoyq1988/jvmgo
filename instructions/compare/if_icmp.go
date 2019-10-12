package compare

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IfICmpEQ struct {
	base.BranchInstruction
}

func (ific *IfICmpEQ) Execute(frame *rtda.Frame) {
	if v1, v2 := icmp(frame); v1 == v2 {
		base.Branch(frame, ific.Offset)
	}
}

type IfICmpNE struct {
	base.BranchInstruction
}

func (ific *IfICmpNE) Execute(frame *rtda.Frame) {
	if v1, v2 := icmp(frame); v1 != v2 {
		base.Branch(frame, ific.Offset)
	}
}

type IfICmpLT struct {
	base.BranchInstruction
}

func (ific *IfICmpLT) Execute(frame *rtda.Frame) {
	if v1, v2 := icmp(frame); v1 < v2 {
		base.Branch(frame, ific.Offset)
	}
}

type IfICmpGE struct {
	base.BranchInstruction
}

func (ific *IfICmpGE) Execute(frame *rtda.Frame) {
	if v1, v2 := icmp(frame); v1 >= v2 {
		base.Branch(frame, ific.Offset)
	}
}

type IfICmpGT struct {
	base.BranchInstruction
}

func (ific *IfICmpGT) Execute(frame *rtda.Frame) {
	if v1, v2 := icmp(frame); v1 > v2 {
		base.Branch(frame, ific.Offset)
	}
}

type IfICmpLE struct {
	base.BranchInstruction
}

func (ific *IfICmpLE) Execute(frame *rtda.Frame) {
	if v1, v2 := icmp(frame); v1 <= v2 {
		base.Branch(frame, ific.Offset)
	}
}

func icmp(frame *rtda.Frame) (int32, int32) {
	v2 := frame.PopInt()
	v1 := frame.PopInt()
	return v1, v2
}
