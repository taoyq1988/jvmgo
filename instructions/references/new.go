package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type New struct {
	base.Index16Instruction
	class *heap.Class
}

func (_new *New) Execute(frame *rtda.Frame) {
	if _new.class == nil {
		cp := frame.GetConstantPool()
		kclass := cp.GetConstantClass(_new.Index)
		_new.class = kclass.GetClass()
	}

	if _new.class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(_new.class)
		return
	}
	ref := _new.class.NewObj()
	frame.PushRef(ref)
}
