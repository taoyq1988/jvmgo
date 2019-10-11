package stack

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type Dup struct {
	base.NoOperandsInstruction
}

func (_ *Dup) Execute(frame *rtda.Frame) {
	val := frame.Pop()
	frame.Push(val)
	frame.Push(val)
}

type DupX1 struct {
	base.NoOperandsInstruction
}

func (_ *DupX1) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	frame.Push(val1)
	frame.Push(val2)
	frame.Push(val1)
}

type DupX2 struct {
	base.NoOperandsInstruction
}

func (_ *DupX2) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	val3 := frame.Pop()
	frame.Push(val1)
	frame.Push(val3)
	frame.Push(val2)
	frame.Push(val1)
}

type Dup2 struct {
	base.NoOperandsInstruction
}

func (dup *Dup2) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	frame.Push(val2)
	frame.Push(val1)
	frame.Push(val2)
	frame.Push(val1)
}

// Duplicate the top one or two operand stack values and insert two or three values down
type Dup2X1 struct {
	base.NoOperandsInstruction
}

func (instr *Dup2X1) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	val3 := frame.Pop()
	frame.Push(val2)
	frame.Push(val1)
	frame.Push(val3)
	frame.Push(val2)
	frame.Push(val1)
}

// Duplicate the top one or two operand stack values and insert two, three, or four values down
type Dup2X2 struct {
	base.NoOperandsInstruction
}

func (instr *Dup2X2) Execute(frame *rtda.Frame) {
	val1 := frame.Pop()
	val2 := frame.Pop()
	val3 := frame.Pop()
	val4 := frame.Pop()
	frame.Push(val2)
	frame.Push(val1)
	frame.Push(val4)
	frame.Push(val3)
	frame.Push(val2)
	frame.Push(val1)
}
