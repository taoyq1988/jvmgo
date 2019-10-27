package loads

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

// xload
type Load struct {
	base.Index8Instruction
	L bool
}

func (load *Load) Execute(frame *rtda.Frame) {
	frame.Load(load.Index, load.L)
}

// xload_n
type LoadN struct {
	base.NoOperandsInstruction
	N uint
	L bool
}

func (load *LoadN) Execute(frame *rtda.Frame) {
	frame.Load(load.N, load.L)
}

// Array load
type IALoad struct {
	base.NoOperandsInstruction
}

func (load *IALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Ints()[index]
		frame.PushInt(val)
	}
}

type LALoad struct {
	base.NoOperandsInstruction
}

func (load *LALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Longs()[index]
		frame.PushLong(val)
	}
}

type FALoad struct {
	base.NoOperandsInstruction
}

func (load *FALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Floats()[index]
		frame.PushFloat(val)
	}
}

type DALoad struct {
	base.NoOperandsInstruction
}

func (load *DALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Doubles()[index]
		frame.PushDouble(val)
	}
}

type AALoad struct {
	base.NoOperandsInstruction
}

func (load *AALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		ref := arrRef.Refs()[index]
		frame.PushRef(ref)
	}
}

type BALoad struct {
	base.NoOperandsInstruction
}

func (load *BALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Bytes()[index]
		frame.PushInt(int32(val))
	}
}

type CALoad struct {
	base.NoOperandsInstruction
}

func (load *CALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Chars()[index]
		frame.PushInt(int32(val))
	}
}

type SALoad struct {
	base.NoOperandsInstruction
}

func (load *SALoad) Execute(frame *rtda.Frame) {
	arrRef, index, ok := _aLoadPop(frame)
	if ok {
		val := arrRef.Shorts()[index]
		frame.PushInt(int32(val))
	}
}

func _aLoadPop(frame *rtda.Frame) (*heap.Object, int32, bool) {
	index := frame.PopInt()
	arrRef := frame.PopRef()
	if arrRef == nil {
		frame.Thread.ThrowNPE()
		return nil, 0, false
	}
	if index < 0 || index >= heap.ArrayLength(arrRef) {
		frame.Thread.ThrowArrayIndexOutOfBoundsException(index)
		return nil, 0, false
	}
	return arrRef, index, true
}
