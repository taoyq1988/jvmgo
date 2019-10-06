package loads

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

// xload
type Load struct {
	base.Index8Instruction
	L bool
}

func (load *Load) Execute(frame *rtda.Frame) {
	frame.Load(load.Index, load.L)
}

// xload_n
type LoadN struct {
	base.NoOperandsInstruction
	N uint
	L bool
}

func (load *LoadN) Execute(frame *rtda.Frame) {
	frame.Load(load.N, load.L)
}

type IALoad struct {
	base.NoOperandsInstruction
}

type LALoad struct {
	base.NoOperandsInstruction
}

type FALoad struct {
	base.NoOperandsInstruction
}

type DALoad struct {
	base.NoOperandsInstruction
}

type AALoad struct {
	base.NoOperandsInstruction
}

type BALoad struct {
	base.NoOperandsInstruction
}

type CALoad struct {
	base.NoOperandsInstruction
}

type SALoad struct {
	base.NoOperandsInstruction
}
