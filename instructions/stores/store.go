package stores

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

// Store store localVar to stack top
type Store struct {
	base.Index8Instruction
	L bool
}

func (store *Store) Execute(frame *rtda.Frame) {
	frame.Store(store.Index, store.L)
}

type StoreX struct {
	base.NoOperandsInstruction
	X uint
	L bool
}

func (store StoreX) Execute(frame *rtda.Frame) {
	frame.Store(store.X, store.L)
}

/**
Array Store
*/
type IAStore struct {
	base.NoOperandsInstruction
}

func (store *IAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Ints()[index] = val
	}
}

type LAStore struct {
	base.NoOperandsInstruction
}

func (store *LAStore) Execute(frame *rtda.Frame) {
	val := frame.PopLong()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Longs()[index] = val
	}
}

type FAStore struct {
	base.NoOperandsInstruction
}

func (store *FAStore) Execute(frame *rtda.Frame) {
	val := frame.PopFloat()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Floats()[index] = val
	}
}

type DAStore struct {
	base.NoOperandsInstruction
}

func (store *DAStore) Execute(frame *rtda.Frame) {
	val := frame.PopDouble()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Doubles()[index] = val
	}
}

type AAStore struct {
	base.NoOperandsInstruction
}

func (store *AAStore) Execute(frame *rtda.Frame) {
	ref := frame.PopRef()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Refs()[index] = ref
	}
}

type BAStore struct {
	base.NoOperandsInstruction
}

func (store *BAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Bytes()[index] = int8(val)
	}
}

type CAStore struct {
	base.NoOperandsInstruction
}

func (store *CAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Chars()[index] = uint16(val)
	}
}

type SAStore struct {
	base.NoOperandsInstruction
}

func (store *SAStore) Execute(frame *rtda.Frame) {
	val := frame.PopInt()
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if _checkArrayAndIndex(frame, arrRef, index) {
		arrRef.Shorts()[index] = int16(val)
	}
}

func _checkArrayAndIndex(frame *rtda.Frame, arrRef *heap.Object, index int32) bool {
	if arrRef == nil {
		frame.Thread.ThrowNEP()
		return false
	}

	if index < 0 || index >= heap.ArrayLength(arrRef) {
		frame.Thread.ThrowArrayIndexOutOfBoundsException(index)
		return false
	}
	return true
}
