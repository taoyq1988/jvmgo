package compare

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type DCmpL struct {
	base.NoOperandsInstruction
}

func (_ *DCmpL) Execute(frame *rtda.Frame) {
	dcmp(frame, false)
}

type DCmpG struct {
	base.NoOperandsInstruction
}

func (_ *DCmpG) Execute(frame *rtda.Frame) {
	dcmp(frame, true)
}

func dcmp(frame *rtda.Frame, isG bool) {
	v1 := frame.PopDouble()
	v2 := frame.PopDouble()
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
