package rtda

import "github.com/taoyq1988/jvmgo/rtda/heap"

type OperandStack struct {
	size  uint
	slots []heap.Slot
}

func newOperandStack(size uint) OperandStack {
	slots := make([]heap.Slot, size)
	return OperandStack{
		size:  0,
		slots: slots,
	}
}

func (stack *OperandStack) PushNull() {
	stack.Push(heap.EmptySlot)
}

func (stack *OperandStack) Push(slot heap.Slot) {
	stack.slots[stack.size] = slot
	stack.size++
}

func (stack *OperandStack) Pop() heap.Slot {
	stack.size--
	top := stack.slots[stack.size]
	stack.slots[stack.size] = heap.EmptySlot // help GC
	return top
}

func (stack *OperandStack) PushInt(val int32) {
	stack.Push(heap.NewIntSlot(val))
}

func (stack *OperandStack) PopInt() int32 {
	return stack.Pop().IntValue()
}

// PushLong long or double need two slots
func (stack *OperandStack) PushLong(val int64) {
	stack.Push(heap.NewLongSlot(val))
	stack.size++
}

func (stack *OperandStack) PopLong() int64 {
	stack.size--
	return stack.Pop().LongValue()
}

func (stack *OperandStack) PushFloat(val float32) {
	stack.Push(heap.NewFloatSlot(val))
}

func (stack *OperandStack) PopFloat() float32 {
	return stack.Pop().FloatValue()
}

func (stack *OperandStack) PushDouble(val float64) {
	stack.Push(heap.NewDoubleSlot(val))
	stack.size++
}

func (stack *OperandStack) PopDouble() float64 {
	stack.size--
	return stack.Pop().DoubleValue()
}

func (stack *OperandStack) PushRef(ref *heap.Object) {
	stack.Push(heap.NewRefSlot(ref))
}

func (stack *OperandStack) PopRef() *heap.Object {
	return stack.Pop().Ref
}

func (stack *OperandStack) PushL(slot heap.Slot, isLongOrDouble bool) {
	stack.Push(slot)
	if isLongOrDouble {
		stack.size++
	}
}

func (stack *OperandStack) PopL(isLongOrDouble bool) heap.Slot {
	if isLongOrDouble {
		stack.size--
	}
	return stack.Pop()
}

func (stack *OperandStack) PopTops(n uint) []heap.Slot {
	start := stack.size - n
	end := stack.size
	stack.size -= n
	return stack.slots[start:end]
}

func (stack *OperandStack) TopRef(n uint) *heap.Object {
	return stack.slots[stack.size-n-1].Ref
}
