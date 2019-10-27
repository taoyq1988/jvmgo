package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type AThrow struct {
	base.NoOperandsInstruction
}

func (aThrow *AThrow) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	ex := frame.PopRef()
	if ex == nil {
		thread.ThrowNPE()
		return
	}

	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC - 1
		handlePC := frame.Method.FindExceptionHandle(ex.Class, pc)
		if handlePC >= 0 {
			frame.ClearStack()
			frame.PushRef(ex)
			frame.NextPC = handlePC
			return
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	thread.HandleUncaughtException(ex)
}
