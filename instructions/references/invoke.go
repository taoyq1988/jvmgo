package references

import (
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
	cp := frame.GetConstantPool()
	k := cp.GetConstant(invoke.Index)
	var method *heap.Method
	switch x := k.(type) {
	case *heap.ConstantMethodRef:
		method = x.GetMethod(false)
	case *heap.ConstantInterfaceMethodRef:
		method = x.GetMethod(false)
	}
	frame.Thread.InvokeMethod(method)
}

// Invoke instance method; dispatch based on class
type InvokeVirtual struct {
	base.Index16Instruction
	kMethodRef    *heap.ConstantMethodRef
	argSlotsCount uint
}

func (invoke *InvokeVirtual) Execute(frame *rtda.Frame) {
	if invoke.kMethodRef == nil {
		cp := frame.GetConstantPool()
		invoke.kMethodRef = cp.GetConstant(invoke.Index).(*heap.ConstantMethodRef)
		invoke.argSlotsCount = invoke.kMethodRef.ParamSlotCount
	}
	ref := frame.TopRef(invoke.argSlotsCount)
	if ref == nil {
		frame.Thread.ThrowNPE()
		return
	}

	method := invoke.kMethodRef.GetVirtualMethod(ref)
	frame.Thread.InvokeMethod(method)
}

// Invoke a class(static) method
type InvokeStatic struct {
	base.Index16Instruction
	method *heap.Method
}

func (invoke *InvokeStatic) Execute(frame *rtda.Frame) {
	if invoke.method == nil {
		cp := frame.GetConstantPool()
		k := cp.GetConstant(invoke.Index)
		switch x := k.(type) {
		case *heap.ConstantMethodRef:
			invoke.method = x.GetMethod(true)
		case *heap.ConstantInterfaceMethodRef:
			invoke.method = x.GetMethod(true)
		}
	}
	class := invoke.method.Class
	if class.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(class)
		return
	}
	frame.Thread.InvokeMethod(invoke.method)
}

// Invoke interface method
type InvokeInterface struct {
	index         uint
	kMethodRef    *heap.ConstantInterfaceMethodRef
	argSlotsCount uint
}

func (invoke *InvokeInterface) FetchOperands(reader *base.CodeReader) {
	invoke.index = uint(reader.ReadUint16())
	reader.ReadUint8() // count
	reader.ReadUint8() // zero
}

func (invoke *InvokeInterface) Execute(frame *rtda.Frame) {
	if invoke.kMethodRef == nil {
		cp := frame.GetConstantPool()
		invoke.kMethodRef = cp.GetConstant(invoke.index).(*heap.ConstantInterfaceMethodRef)
		invoke.argSlotsCount = invoke.kMethodRef.ParamSlotCount
	}
	ref := frame.TopRef(invoke.argSlotsCount)
	if ref == nil {
		frame.Thread.ThrowNPE()
		return
	}

	method := invoke.kMethodRef.FindInterfaceMethod(ref)
	frame.Thread.InvokeMethod(method)
}
