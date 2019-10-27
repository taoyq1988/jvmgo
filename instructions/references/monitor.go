package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type MonitorEnter struct {
	base.NoOperandsInstruction
}

func (monitor *MonitorEnter) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	ref := frame.PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor.Enter(thread)
	}
}

type MonitorExit struct {
	base.NoOperandsInstruction
}

func (monitor *MonitorExit) Execute(frame *rtda.Frame) {
	thread := frame.Thread
	ref := frame.PopRef()
	if ref == nil {
		frame.RevertNextPC()
		thread.ThrowNPE()
	} else {
		ref.Monitor.Exit(thread)
	}
}
