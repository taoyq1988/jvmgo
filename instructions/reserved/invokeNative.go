package reserved

import (
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

// Invoke native method
type InvokeNative struct{ base.NoOperandsInstruction }

func (instr *InvokeNative) Execute(frame *rtda.Frame) {
	nativeMethod := frame.Method.GetNativeMethod().(func(*rtda.Frame))
	nativeMethod(frame)
}
