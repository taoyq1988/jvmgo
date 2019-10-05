package constants

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (nop *NOP) Execute(frame *rtda.Frame) {

}
