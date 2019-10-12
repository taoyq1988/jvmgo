package compare

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type LCmp struct {
	base.NoOperandsInstruction
}

func (_ *LCmp) Execute(frame *rtda.Frame) {
	v1 := frame.PopLong()
	v2 := frame.PopLong()
	if v2 > v1 {
		frame.PushInt(1)
	} else if v2 == v1 {
		frame.PushInt(0)
	} else {
		frame.PushInt(-1)
	}
}
