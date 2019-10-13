package control

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type Jsr struct {
	base.BranchInstruction
}

func (_ *Jsr) Execute(frame *rtda.Frame) {
	//todo
	panic("jsr")
}

// Return from subroutine
type Ret struct{ base.Index8Instruction }

func (instr *Ret) Execute(frame *rtda.Frame) {
	panic("ret")
}

// Jump subroutine (wide index)
type JsrW struct {
	offset int
}

func (instr *JsrW) FetchOperands(reader *base.CodeReader) {
	instr.offset = int(reader.ReadInt32())
}
func (instr *JsrW) Execute(frame *rtda.Frame) {
	panic("todo")
}
