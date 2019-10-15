package base

import (
	"github.com/taoyq1988/jvmgo/rtda"
)

type Instruction interface {
	FetchOperands(reader *CodeReader)
	Execute(frame *rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (_ *NoOperandsInstruction) FetchOperands(reader *CodeReader) {

}

type BranchInstruction struct {
	Offset int
}

func (instr *BranchInstruction) FetchOperands(reader *CodeReader) {
	instr.Offset = int(reader.ReadInt16())
}

type Index8Instruction struct {
	Index uint
}

func (instr *Index8Instruction) FetchOperands(reader *CodeReader) {
	instr.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index uint
}

func (instr *Index16Instruction) FetchOperands(reader *CodeReader) {
	instr.Index = uint(reader.ReadUint16())
}

func Branch(frame *rtda.Frame, offset int) {
	frame.NextPC = frame.Thread.PC + offset
}
