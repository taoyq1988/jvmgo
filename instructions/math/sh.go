package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IShl struct {
	base.NoOperandsInstruction
}

func (_ IShl) Execute(frame *rtda.Frame) {
	p := frame.PopInt()
	val := frame.PopInt()
	s := uint32(p) & 0x1f
	frame.PushInt(val << s)
}

type LShl struct {
	base.NoOperandsInstruction
}

func (_ LShl) Execute(frame *rtda.Frame) {
	p := frame.PopInt()
	val := frame.PopLong()
	s := uint32(p) & 0x3f
	frame.PushLong(val << s)
}

type IShr struct {
	base.NoOperandsInstruction
}

func (_ IShr) Execute(frame *rtda.Frame) {
	p := frame.PopInt()
	val := frame.PopInt()
	s := uint32(p) & 0x1f
	frame.PushInt(val >> s)
}

type LShr struct {
	base.NoOperandsInstruction
}

func (_ LShr) Execute(frame *rtda.Frame) {
	p := frame.PopInt()
	val := frame.PopLong()
	s := uint32(p) & 0x3f
	frame.PushLong(val >> s)
}

type IUShr struct {
	base.NoOperandsInstruction
}

func (_ IUShr) Execute(frame *rtda.Frame) {
	p := frame.PopInt()
	val := frame.PopInt()
	s := uint32(p) & 0x1f
	frame.PushInt(int32(uint32(val) >> s))
}

type LUShr struct {
	base.NoOperandsInstruction
}

func (_ LUShr) Execute(frame *rtda.Frame) {
	p := frame.PopInt()
	val := frame.PopLong()
	s := uint32(p) & 0x3f
	frame.PushLong(int64(uint64(val) >> s))
}
