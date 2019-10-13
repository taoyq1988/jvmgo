package rtda

import "fmt"

// Stack jvm stack
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (stack *Stack) isEmpty() bool {
	return stack._top == nil
}

func (stack *Stack) push(frame *Frame) {
	if stack.size >= stack.maxSize {
		//todo throw stack over flow
		panic("StackOverFlow")
	}

	if stack._top != nil {
		frame.lower = stack._top
	}
	stack._top = frame
	stack.size++
}

func (stack *Stack) pop() *Frame {
	if stack._top == nil {
		panic("stack is empty")
	}
	r := stack._top
	stack._top = r.lower
	r.lower = nil
	stack.size--
	return r
}

func (stack *Stack) clear() {
	for stack._top != nil {
		stack.pop()
	}
}

func (stack *Stack) top() *Frame {
	if stack._top == nil {
		panic("stack is empty")
	}
	return stack._top
}

func (stack *Stack) topN(n uint) *Frame {
	if stack.size < n {
		panic(fmt.Sprintf("jvm stack size:%v n:%v", stack.size, n))
	}

	frame := stack._top
	for n > 0 {
		frame = frame.lower
		n--
	}

	return frame
}
