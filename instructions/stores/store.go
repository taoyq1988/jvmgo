package stores

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

// Store store localVar to stack top
type Store struct {
	base.Index8Instruction
	L bool
}

func (store *Store) Execute(frame *rtda.Frame) {
	frame.Store(store.Index, store.L)
}

type StoreX struct {
	base.NoOperandsInstruction
	X uint
	L bool
}

func (store StoreX) Execute(frame *rtda.Frame) {
	frame.Store(store.X, store.L)
}

type IAStore struct {
	base.NoOperandsInstruction
}

type LAStore struct {
	base.NoOperandsInstruction
}

type FAStore struct {
	base.NoOperandsInstruction
}

type DAStore struct {
	base.NoOperandsInstruction
}

type AAStore struct {
	base.NoOperandsInstruction
}

type BAStore struct {
	base.NoOperandsInstruction
}

type CAStore struct {
	base.NoOperandsInstruction
}

type SAStore struct {
	base.NoOperandsInstruction
}
