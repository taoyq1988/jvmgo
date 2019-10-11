package constants

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type LDC struct {
	base.Index8Instruction
}

func (ldc *LDC) Execute(frame *rtda.Frame) {
	_ldc(frame, ldc.Index)
}

type LDC_W struct {
	base.Index16Instruction
}

func (ldc *LDC_W) Execute(frame *rtda.Frame) {
	_ldc(frame, ldc.Index)
}

//TODO
func _ldc(frame *rtda.Frame, index uint) {
	frame.GetConstantPool()
}
