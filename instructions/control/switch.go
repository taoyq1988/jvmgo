package control

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

type TableSwitch struct {
	DefaultOffset int32
	low           int32
	high          int32
	JumpOffsets   []int32
}

func (tSwitch *TableSwitch) FetchOperands(reader *base.CodeReader) {
	reader.SkipPadding()
	tSwitch.DefaultOffset = reader.ReadInt32()
	tSwitch.low = reader.ReadInt32()
	tSwitch.high = reader.ReadInt32()
	jumpOffsetsCount := tSwitch.high - tSwitch.low + 1
	tSwitch.JumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (tSwitch *TableSwitch) Execute(frame *rtda.Frame) {
	index := frame.PopInt()

	var offset int
	if index >= tSwitch.low && index <= tSwitch.high {
		offset = int(tSwitch.JumpOffsets[index-tSwitch.low])
	} else {
		offset = int(tSwitch.DefaultOffset)
	}

	base.Branch(frame, offset)
}

type LookupSwitch struct {
	DefaultOffset int32
	npairs        int32
	MatchOffsets  []int32
}

func (instr *LookupSwitch) FetchOperands(reader *base.CodeReader) {
	reader.SkipPadding()
	instr.DefaultOffset = reader.ReadInt32()
	instr.npairs = reader.ReadInt32()
	instr.MatchOffsets = reader.ReadInt32s(instr.npairs * 2)
}

func (instr *LookupSwitch) Execute(frame *rtda.Frame) {
	key := frame.PopInt()
	for i := int32(0); i < instr.npairs*2; i += 2 {
		if instr.MatchOffsets[i] == key {
			offset := instr.MatchOffsets[i+1]
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(instr.DefaultOffset))
}
