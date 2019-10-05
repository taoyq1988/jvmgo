package heap

import "math"

var EmptySlot = Slot{0, nil}

type Slot struct {
	Val int64
	Ref *Object
}

func NewIntSlot(n int32) Slot {
	return Slot{Val: int64(n)}
}

func NewLongSlot(n int64) Slot {
	return Slot{Val: n}
}

func NewFloatSlot(f float32) Slot {
	return Slot{Val: int64(math.Float32bits(f))}
}

func NewDoubleSlot(f float64) Slot {
	return Slot{Val: int64(math.Float64bits(f))}
}

func NewRefSlot(ref *Object) Slot {
	return Slot{Ref: ref}
}

func (slot Slot) IntValue() int32 {
	return int32(slot.Val)
}

func (slot Slot) LongValue() int64 {
	return slot.Val
}

func (slot Slot) FloatValue() float32 {
	return math.Float32frombits(uint32(slot.Val))
}

func (slot Slot) DoubleValue() float64 {
	return math.Float64frombits(uint64(slot.Val))
}
