package rtda

type Frame struct {
	lower *Frame
	LocalVars
	OperandStack
}
