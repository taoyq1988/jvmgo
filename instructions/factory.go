package instructions

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	. "github.com/taoyq1988/jvmgo/instructions/constants"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

// No Operand Instructions(singleton)
var (
	nop         = &NOP{}
	aconst_null = &Const{K: heap.EmptySlot}
	iconst_m1   = &Const{K: heap.NewIntSlot(-1)}
	iconst_0    = &Const{K: heap.NewIntSlot(0)}
	iconst_1    = &Const{K: heap.NewIntSlot(1)}
	iconst_2    = &Const{K: heap.NewIntSlot(2)}
	iconst_3    = &Const{K: heap.NewIntSlot(3)}
	iconst_4    = &Const{K: heap.NewIntSlot(4)}
	iconst_5    = &Const{K: heap.NewIntSlot(5)}
	lconst_0    = &Const{K: heap.NewLongSlot(0), L: true}
	lconst_1    = &Const{K: heap.NewLongSlot(1), L: true}
	fconst_0    = &Const{K: heap.NewFloatSlot(0.0)}
	fconst_1    = &Const{K: heap.NewFloatSlot(1.0)}
	fconst_2    = &Const{K: heap.NewFloatSlot(2.0)}
	dconst_0    = &Const{K: heap.NewDoubleSlot(0.0), L: true}
	dconst_1    = &Const{K: heap.NewDoubleSlot(1.0), L: true}

)

func newInstruction(opcode byte) base.Instruction {
	switch opcode {
	case OpNop:
		return nop
	case OpAConstNull:
		return aconst_null
	case OpIConstM1:
		return iconst_m1
	case OpIConst0:
		return iconst_0
	case OpIConst1:
		return iconst_1
	case OpIConst2:
		return iconst_2
	case OpIConst3:
		return iconst_3
	case OpIConst4:
		return iconst_4
	case OpIConst5:
		return iconst_5
	case OpLConst0:
		return lconst_0
	case OpLConst1:
		return lconst_1
	case OpFConst0:
		return fconst_0
	case OpFConst1:
		return fconst_1
	case OpFConst2:
		return fconst_2
	case OpDConst0:
		return dconst_0
	case OpDConst1:
		return dconst_1
	case OpBIPush:
		return &BIPush{}
	case OpSIPush:
		return &SIPush{}
	default:
		return nil
	}
}
