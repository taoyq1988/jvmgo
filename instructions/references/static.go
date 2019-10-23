package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

// set static field in class
type PutStatic struct {
	base.Index16Instruction
	field *heap.Field
}

func (static *PutStatic) Execute(frame *rtda.Frame) {
	if static.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(static.Index)
		static.field = fieldRef.GetField(true)
	}

	class := static.field.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(class)
	}

	val := frame.PopL(static.field.IsLongOrDouble)
	static.field.PutStaticValue(val)
}

// fetch field from object
type GetStatic struct {
	base.Index16Instruction
	field *heap.Field
}

func (static *GetStatic) Execute(frame *rtda.Frame) {
	if static.field == nil {
		cp := frame.GetConstantPool()
		fieldRef := cp.GetConstantFieldRef(static.Index)
		static.field = fieldRef.GetField(true)
	}

	class := static.field.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(class)
		return
	}

	val := static.field.GetStaticValue()
	frame.PushL(val, static.field.IsLongOrDouble)
}
