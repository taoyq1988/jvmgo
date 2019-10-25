package references

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

type NewArray struct {
	aType uint8
}

func (array *NewArray) FetchOperands(reader *base.CodeReader) {
	array.aType = reader.ReadUint8()
}

func (array *NewArray) Execute(frame *rtda.Frame) {
	count := frame.PopInt()
	if count < 0 {
		frame.Thread.ThrowNegativeArraySizeException()
		return
	}

	arr := heap.NewPrimitiveArray(array.aType, uint(count))
	frame.PushRef(arr)
}


type ANewArray struct {
	base.Index16Instruction
}

func (array *ANewArray) Execute(frame *rtda.Frame) {
	cp := frame.GetConstantPool()
	kClass := cp.GetConstantClass(array.Index)
	componentClass := kClass.GetClass()
	if componentClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(componentClass)
		return
	}
	count := frame.PopInt()
	if count < 0 {
		frame.Thread.ThrowNegativeArraySizeException()
	} else {
		arr := heap.NewRefArray(componentClass, uint(count))
		frame.PushRef(arr)
	}
}


type MultiANewArray struct {
	index uint16
	dimensions uint8
}

func (array *MultiANewArray) FetchOperands(reader *base.CodeReader) {
	array.index = reader.ReadUint16()
	array.dimensions = reader.ReadUint8()
}

func (array *MultiANewArray) Execute(frame *rtda.Frame) {
	cp := frame.GetConstantPool()
	arrClass := cp.GetConstantClass(uint(array.index)).GetClass()
	counts := frame.PopTops(uint(array.dimensions))
	if _checkCounts(counts) {
		arr := _newMultiArray(counts, arrClass)
		frame.PushRef(arr)
	} else {
		frame.Thread.ThrowNegativeArraySizeException()
	}
}

func _checkCounts(counts []heap.Slot) bool {
	for _, c := range counts {
		if c.IntValue() < 0 {
			return false
		}
	}
	return true
}

func _newMultiArray(counts []heap.Slot, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0].IntValue())
	arr := heap.NewArray(arrClass, count)
	if len(counts) > 1 {
		objs := arr.Refs()
		for i := range objs {
			objs[i] = _newMultiArray(counts[1:], arrClass.ComponentClass())
		}
	}
	return arr
}


type ArrayLength struct {
	base.NoOperandsInstruction
}

func (al *ArrayLength) Execute(frame *rtda.Frame) {
	arrRef := frame.PopRef()
	if arrRef == nil {
		frame.Thread.ThrowNEP()
		return
	}
	arrLen := heap.ArrayLength(arrRef)
	frame.PushInt(arrLen)
}
