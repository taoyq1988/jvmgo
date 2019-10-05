package rtda

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
