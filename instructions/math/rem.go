package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"math"
)

type IRem struct {
	base.NoOperandsInstruction
}

func (_ *IRem) Execute(frame *rtda.Frame) {
	val1 := frame.PopInt()
	val2 := frame.PopInt()
	if val1 == 0 {
		//todo
	}
	frame.PushInt(val2 % val1)
}

type LRem struct {
	base.NoOperandsInstruction
}

func (_ *LRem) Execute(frame *rtda.Frame) {
	val1 := frame.PopLong()
	val2 := frame.PopLong()
	if val1 == 0 {
		//todo
	}
	frame.PushLong(val2 % val1)
}

type FRem struct {
	base.NoOperandsInstruction
}

func (_ *FRem) Execute(frame *rtda.Frame) {
	val1 := frame.PopFloat()
	val2 := frame.PopFloat()
	if val1 == 0 {
		//todo
	}
	frame.PushFloat(float32(math.Mod(float64(val2), float64(val1))))
}

type DRem struct {
	base.NoOperandsInstruction
}

func (_ *DRem) Execute(frame *rtda.Frame) {
	val1 := frame.PopDouble()
	val2 := frame.PopDouble()
	if val1 == 0 {
		//todo
	}
	frame.PushDouble(math.Mod(val2, val1))
}
