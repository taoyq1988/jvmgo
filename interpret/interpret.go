package interpret

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/instructions"
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

const isLog = false

func Interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	defer catchError(frame)
	loop(thread)
}

func catchError(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Printf("[DEBUG] frame method %s, method descriptor %s, method class %s\n", frame.Method.Name, frame.Method.Descriptor, frame.Method.Class.Name)
		fmt.Println("[DEBUG] localvars", frame.LocalVars)
		fmt.Println("[DEBUG] operandstack", frame.OperandStack)
		panic("error")
	}
}

func logInstruction(frame *rtda.Frame, inst base.Instruction) {
	fmt.Printf("[execute] method %s, class %s, inst %T %v, pc %d\n", frame.Method.Name, frame.Method.Class.Name, inst, inst, frame.Thread.PC)
	fmt.Println("[execute] localvars", frame.LocalVars)
	fmt.Println("[execute] operandstack", frame.OperandStack)
	fmt.Println()
}

func loop(thread *rtda.Thread) {
	reader := &base.CodeReader{}
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC
		thread.PC = pc
		reader.Reset(frame.Method.Code, thread.PC)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.NextPC = reader.PC()
		if isLog {
			logInstruction(frame, inst)
		}
		inst.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}
