package compare

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type FCmpL struct {
	base.NoOperandsInstruction
}

func (_ *FCmpL) Execute(frame *rtda.Frame) {
	fcmp(frame, false)
}

type FCmpG struct {
	base.NoOperandsInstruction
}

func (_ *FCmpG) Execute(frame *rtda.Frame) {
	fcmp(frame, true)
}

func fcmp(frame *rtda.Frame, isG bool) {
	v1 := frame.PopFloat()
	v2 := frame.PopFloat()
	if v2 > v1 {
		frame.PushInt(1)
	} else if v2 == v1 {
		frame.PushInt(0)
	} else if v2 < v1 {
		frame.PushInt(-1)
	} else if isG {
		frame.PushInt(1)
	} else {
		frame.PushInt(-1)
	}
}
