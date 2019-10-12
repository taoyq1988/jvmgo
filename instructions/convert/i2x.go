package convert

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type I2L struct {
	base.NoOperandsInstruction
}

func (_ *I2L) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	frame.PushLong(int64(i))
}

type I2F struct {
	base.NoOperandsInstruction
}

func (_ *I2F) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	frame.PushFloat(float32(i))
}

type I2D struct {
	base.NoOperandsInstruction
}

func (_ *I2D) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	frame.PushDouble(float64(i))
}

type I2B struct {
	base.NoOperandsInstruction
}

func (_ *I2B) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	frame.PushInt(int32(int8(i)))
}

type I2C struct {
	base.NoOperandsInstruction
}

func (_ *I2C) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	frame.PushInt(int32(uint16(i)))
}

type I2S struct {
	base.NoOperandsInstruction
}

func (_ *I2S) Execute(frame *rtda.Frame) {
	i := frame.PopInt()
	frame.PushInt(int32(int16(i)))
}
