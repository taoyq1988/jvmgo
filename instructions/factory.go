package instructions

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	. "github.com/taoyq1988/jvmgo/instructions/compare"
	. "github.com/taoyq1988/jvmgo/instructions/constants"
	. "github.com/taoyq1988/jvmgo/instructions/control"
	. "github.com/taoyq1988/jvmgo/instructions/convert"
	. "github.com/taoyq1988/jvmgo/instructions/extend"
	. "github.com/taoyq1988/jvmgo/instructions/loads"
	. "github.com/taoyq1988/jvmgo/instructions/math"
	. "github.com/taoyq1988/jvmgo/instructions/options"
	. "github.com/taoyq1988/jvmgo/instructions/stack"
	. "github.com/taoyq1988/jvmgo/instructions/stores"
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
	iload_0     = &LoadN{N: 0}
	iload_1     = &LoadN{N: 1}
	iload_2     = &LoadN{N: 2}
	iload_3     = &LoadN{N: 3}
	lload_0     = &LoadN{N: 0, L: true}
	lload_1     = &LoadN{N: 1, L: true}
	lload_2     = &LoadN{N: 2, L: true}
	lload_3     = &LoadN{N: 3, L: true}
	fload_0     = &LoadN{N: 0}
	fload_1     = &LoadN{N: 1}
	fload_2     = &LoadN{N: 2}
	fload_3     = &LoadN{N: 3}
	dload_0     = &LoadN{N: 0, L: true}
	dload_1     = &LoadN{N: 1, L: true}
	dload_2     = &LoadN{N: 2, L: true}
	dload_3     = &LoadN{N: 3, L: true}
	aload_0     = &LoadN{N: 0}
	aload_1     = &LoadN{N: 1}
	aload_2     = &LoadN{N: 2}
	aload_3     = &LoadN{N: 3}
	istore_0    = &StoreX{X: 0}
	istore_1    = &StoreX{X: 1}
	istore_2    = &StoreX{X: 2}
	istore_3    = &StoreX{X: 3}
	lstore_0    = &StoreX{X: 0, L: true}
	lstore_1    = &StoreX{X: 1, L: true}
	lstore_2    = &StoreX{X: 2, L: true}
	lstore_3    = &StoreX{X: 3, L: true}
	fstore_0    = &StoreX{X: 0}
	fstore_1    = &StoreX{X: 1}
	fstore_2    = &StoreX{X: 2}
	fstore_3    = &StoreX{X: 3}
	dstore_0    = &StoreX{X: 0, L: true}
	dstore_1    = &StoreX{X: 1, L: true}
	dstore_2    = &StoreX{X: 2, L: true}
	dstore_3    = &StoreX{X: 3, L: true}
	astore_0    = &StoreX{X: 0}
	astore_1    = &StoreX{X: 1}
	astore_2    = &StoreX{X: 2}
	astore_3    = &StoreX{X: 3}
	pop         = &Pop{}
	pop2        = &Pop2{}
	dup         = &Dup{}
	dupx1       = &DupX1{}
	dupx2       = &DupX2{}
	dup2        = &Dup2{}
	dup2x1      = &Dup2X1{}
	dup2x2      = &Dup2X2{}
	swap        = &Swap{}
	iadd        = &IAdd{}
	ladd        = &LAdd{}
	fadd        = &FAdd{}
	dadd        = &DAdd{}
	isub        = &ISub{}
	lsub        = &LSub{}
	fsub        = &FSub{}
	dsub        = &DSub{}
	imul        = &IMul{}
	lmul        = &LMul{}
	fmul        = &FMul{}
	dmul        = &DMul{}
	idiv        = &IDiv{}
	ldiv        = &LDiv{}
	fdiv        = &FDiv{}
	ddiv        = &DDiv{}
	irem        = &IRem{}
	lrem        = &LRem{}
	frem        = &FRem{}
	drem        = &DRem{}
	ineg        = &INeg{}
	lneg        = &LNeg{}
	fneg        = &FNeg{}
	dneg        = &DNeg{}
	ishl        = &IShl{}
	lshl        = &LShl{}
	ishr        = &IShr{}
	lshr        = &LShr{}
	iushr       = &IUShr{}
	lushr       = &LUShr{}
	iand        = &IAnd{}
	land        = &LAnd{}
	ior         = &IOr{}
	lor         = &LOr{}
	ixor        = &IXor{}
	lxor        = &LXor{}
	i2l         = &I2L{}
	i2f         = &I2F{}
	i2d         = &I2D{}
	l2i         = &L2I{}
	l2f         = &L2F{}
	l2d         = &L2D{}
	f2i         = &F2I{}
	f2l         = &F2L{}
	f2d         = &F2D{}
	d2i         = &D2I{}
	d2l         = &D2L{}
	d2f         = &D2F{}
	i2b         = &I2B{}
	i2c         = &I2C{}
	i2S         = &I2S{}
	lcmp        = &LCmp{}
	fcmpl       = &FCmpL{}
	fcmpg       = &FCmpG{}
	dcmpl       = &DCmpL{}
	dcmpg       = &DCmpG{}
	ireturn     = &IReturn{}
	lreturn     = &LReturn{}
	freturn     = &FReturn{}
	dreturn     = &DReturn{}
	areturn     = &AReturn{}
	_return     = &Return{}
)

func NewInstruction(opcode byte) base.Instruction {
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
	//case OpLDC: //TODO
	//case OpLDCw: //TODO
	//case OpLDC2w: //TODO
	case OpILoad:
		return &Load{}
	case OpLLoad:
		return &Load{L: true}
	case OpFLoad:
		return &Load{}
	case OpDLoad:
		return &Load{L: true}
	case OpALoad:
		return &Load{}
	case OpILoad0:
		return iload_0
	case OpILoad1:
		return iload_1
	case OpILoad2:
		return iload_2
	case OpILoad3:
		return iload_3
	case OpLLoad0:
		return lload_0
	case OpLLoad1:
		return lload_1
	case OpLLoad2:
		return lload_2
	case OpLLoad3:
		return lload_3
	case OpFLoad0:
		return fload_0
	case OpFLoad1:
		return fload_1
	case OpFLoad2:
		return fload_2
	case OpFLoad3:
		return fload_3
	case OpDLoad0:
		return dload_0
	case OpDLoad1:
		return dload_1
	case OpDLoad2:
		return dload_2
	case OpDLoad3:
		return dload_3
	case OpALoad0:
		return aload_0
	case OpALoad1:
		return aload_1
	case OpALoad2:
		return aload_2
	case OpALoad3:
		return aload_3
	//case OpIALoad: //TODO
	//case OpLALoad: //TODO
	//case OpFALoad: //todo
	//case OpDALoad: //todo
	//case OpAALoad: //todo
	//case OpBALoad: //todo
	//case OpCALoad: //todo
	//case OpSALoad: //todo
	case OpIStore:
		return &Store{}
	case OpLStore:
		return &Store{L: true}
	case OpFStore:
		return &Store{}
	case OpDStore:
		return &Store{L: true}
	case OpAStore:
		return &Store{}
	case OpIStore0:
		return istore_0
	case OpIStore1:
		return istore_1
	case OpIStore2:
		return istore_2
	case OpIStore3:
		return istore_3
	case OpLStore0:
		return lstore_0
	case OpLStore1:
		return lstore_1
	case OpLStore2:
		return lstore_2
	case OpLStore3:
		return lstore_3
	case OpFStore0:
		return fstore_0
	case OpFStore1:
		return fstore_1
	case OpFStore2:
		return fstore_2
	case OpFStore3:
		return fstore_3
	case OpDStore0:
		return dstore_0
	case OpDStore1:
		return dstore_1
	case OpDStore2:
		return dstore_2
	case OpDStore3:
		return dstore_3
	case OpAStore0:
		return astore_0
	case OpAStore1:
		return astore_1
	case OpAStore2:
		return astore_2
	case OpAStore3:
		return astore_3
	//case OpIAStore:
	//case OpLAStore:
	//case OpFAStore:
	//case OpDAStore:
	//case OpAAStore:
	//case OpBAStore:
	//case OpCAStore:
	//case OpSAStore:
	case OpPop:
		return pop
	case OpPop2:
		return pop2
	case OpDup:
		return dup
	case OpDupX1:
		return dupx1
	case OpDupX2:
		return dupx2
	case OpDup2:
		return dup2
	case OpDup2X1:
		return dup2x1
	case OpDup2X2:
		return dup2x2
	case OpSwap:
		return swap
	case OpIAdd:
		return iadd
	case OpLAdd:
		return ladd
	case OpFAdd:
		return fadd
	case OpDAdd:
		return dadd
	case OpISub:
		return isub
	case OpLSub:
		return lsub
	case OpFSub:
		return fsub
	case OpDSub:
		return dsub
	case OpIMul:
		return imul
	case OpLMul:
		return lmul
	case OpFMul:
		return fmul
	case OpDMul:
		return dmul
	case OpIDiv:
		return idiv
	case OpLDiv:
		return ldiv
	case OpFDiv:
		return fdiv
	case OpDDiv:
		return ddiv
	case OpIRem:
		return irem
	case OpLRem:
		return lrem
	case OpFRem:
		return frem
	case OpDRem:
		return drem
	case OpINeg:
		return ineg
	case OpLNeg:
		return lneg
	case OpFNeg:
		return fneg
	case OpDNeg:
		return dneg
	case OpIShl:
		return ishl
	case OpLShl:
		return lshl
	case OpIShr:
		return ishr
	case OpLShr:
		return lshr
	case OpIUshr:
		return iushr
	case OpLUshr:
		return lushr
	case OpIAnd:
		return iand
	case OpLAnd:
		return land
	case OpIOr:
		return ior
	case OpLOr:
		return lor
	case OpIXor:
		return ixor
	case OpLXor:
		return lxor
	case OpIInc:
		return &IInc{}
	case OpI2L:
		return i2l
	case OpI2F:
		return i2f
	case OpI2D:
		return i2d
	case OpL2I:
		return l2i
	case OpL2F:
		return l2f
	case OpL2D:
		return l2d
	case OpF2I:
		return f2i
	case OpF2L:
		return f2l
	case OpF2D:
		return f2d
	case OpD2I:
		return d2i
	case OpD2L:
		return d2l
	case OpD2F:
		return d2f
	case OpI2B:
		return i2b
	case OpI2C:
		return i2c
	case OpI2S:
		return i2S
	case OpLCmp:
		return lcmp
	case OpFCmpL:
		return fcmpl
	case OpFCmpG:
		return fcmpg
	case OpDCmpL:
		return dcmpl
	case OpDCmpG:
		return dcmpg
	case OpIfEQ:
		return &IfEQ{}
	case OpIfNE:
		return &IfNE{}
	case OpIfLT:
		return &IfLT{}
	case OpIfGE:
		return &IfGE{}
	case OpIfGT:
		return &IfGT{}
	case OpIfLE:
		return &IfLE{}
	case OpIfICmpEQ:
		return &IfICmpEQ{}
	case OpIfICmpNE:
		return &IfICmpNE{}
	case OpIfICmpLT:
		return &IfICmpLT{}
	case OpIfICmpGE:
		return &IfICmpGE{}
	case OpIfICmpGT:
		return &IfICmpGT{}
	case OpIfICmpLE:
		return &IfICmpLE{}
	case OpIfACmpEQ:
		return &IfACmpEQ{}
	case OpIfACmpNE:
		return &IfACmpNE{}
	case OpGoto:
		return &Goto{}
	case OpJSR:
		return &Jsr{}
	case OpRET:
		return &Ret{}
	case OpTableSwitch:
		return &TableSwitch{}
	case OpLookupSwitch:
		return &LookupSwitch{}
	case OpIReturn:
		return ireturn
	case OpLReturn:
		return lreturn
	case OpFReturn:
		return freturn
	case OpDReturn:
		return dreturn
	case OpAReturn:
		return areturn
	case OpReturn:
		return _return
	case OpWide:
		return &Wide{}
	case OpIfNull:
		return &IfNull{}
	case OpIfNonNull:
		return &IfNonNull{}
	case OpGotoW:
		return &GotoW{}
	case OpJSRw:
		return &JsrW{}
	//case OpBreakpoint: //todo

	default:
		panic("invalid opcode")
	}
}
