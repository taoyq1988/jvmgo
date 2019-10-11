package math

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type IInc struct {
	Index uint
	Const int32
}

func (iinc *IInc) FetchOperands(reader *base.CodeReader) {
	iinc.Index = uint(reader.ReadInt8())
	iinc.Const = int32(reader.ReadInt8())
}

func (iinc *IInc) Execute(frame *rtda.Frame) {
	val := frame.GetIntVar(iinc.Index)
	val += iinc.Const
	frame.SetIntVar(iinc.Index, val)
}
