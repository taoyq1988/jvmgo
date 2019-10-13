package extended

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/instructions/control"
	"github.com/taoyq1988/jvmgo/instructions/loads"
	"github.com/taoyq1988/jvmgo/instructions/math"
	. "github.com/taoyq1988/jvmgo/instructions/options"
	"github.com/taoyq1988/jvmgo/instructions/stores"
	"github.com/taoyq1988/jvmgo/rtda"
)

// Extend local variable index by additional bytes
type Wide struct {
	modifiedInstruction base.Instruction
}

func (instr *Wide) FetchOperands(reader *base.CodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case OpILoad, OpFLoad, OpALoad:
		inst := &loads.Load{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case OpLLoad, OpDLoad:
		inst := &loads.Load{L: true}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case OpIStore, OpFStore, OpAStore:
		inst := &stores.Store{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case OpLStore, OpDStore:
		inst := &stores.Store{L: true}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case OpRET:
		inst := &control.Ret{}
		inst.Index = uint(reader.ReadUint16())
		instr.modifiedInstruction = inst
	case OpIInc:
		inst := &math.IInc{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadInt16())
		instr.modifiedInstruction = inst
	}
}

func (instr *Wide) Execute(frame *rtda.Frame) {
	instr.modifiedInstruction.Execute(frame)
}
