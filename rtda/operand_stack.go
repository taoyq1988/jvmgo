package rtda

import "github.com/taoyq1988/jvmgo/rtda/heap"

type OperandStack struct {
	size  int
	slots []Slot
}

func newOperandStack(size uint) OperandStack {
	slots := make([]Slot, size)
	return OperandStack{
		size:  0,
		slots: slots,
	}
}

func (stack *OperandStack) PushNull() {
	stack.Push(EmptySlot)
}

func (stack *OperandStack) Push(slot Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}

func (stack *OperandStack) Pop() Slot {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = EmptySlot // help GC
	return top
}

func (stack *OperandStack) PushInt(val int32) {
	stack.Push(heap.NewIntSlot(val))
}

func (stack *OperandStack) PopInt() int32 {
	return stack.Pop().IntValue()
}
