package constants

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type BIPush struct {
	Val int32
}

func (push *BIPush) FetchOperands(reader *base.CodeReader) {
	push.Val = int32(reader.ReadInt8())
}

func (push *BIPush) Execute(frame *rtda.Frame) {
	frame.PushInt(push.Val)
}

type SIPush struct {
	Val int32
}

func (push *SIPush) FetchOperands(reader *base.CodeReader) {
	push.Val = int32(reader.ReadInt16())
}

func (push *SIPush) Execute(frame *rtda.Frame) {
	frame.PushInt(push.Val)
}
