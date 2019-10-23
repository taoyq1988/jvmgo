package references

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

// Invoke instance method;
// special handling for superclass, private, and instance initialization method invocations
type InvokeSpecial struct {
	base.Index16Instruction
}

func (invoke *InvokeSpecial) Execute(frame *rtda.Frame) {
	//todo mock empty constructor
	frame.OperandStack.PopRef()
}


// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	base.Index16Instruction
}

func (invoke *InvokeVirtual) Execute(frame *rtda.Frame) {
	//todo mock
	cp := frame.GetConstantPool()
	methodRef := cp.GetConstant(invoke.Index).(*heap.ConstantMethodRef)
	if methodRef.Name() == "println" {
		switch methodRef.Descriptor() {
		case "()V": fmt.Println("**")
		case "(B)V":fmt.Println("**",frame.PopInt())
		case "(I)V":fmt.Println("**", frame.PopInt())
		case "(J)V":fmt.Println("**", frame.PopLong())
		case "(Z)V":fmt.Println("**", frame.PopInt() != 0)
		default:
			panic("** panic")
		}
		frame.PopRef()
	}
}
