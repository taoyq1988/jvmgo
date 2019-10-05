package rtda

// Stack jvm stack
type Stack struct {
	maxSize uint
	size    uint
	top     *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
