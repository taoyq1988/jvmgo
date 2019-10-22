package interpret

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/instructions"
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
)

func Interpret(methodInfo *classfile.MemberInfo) {
}

func catchError(frame *rtda.Frame) {
	if r := recover(); r != nil {
		fmt.Println("localvars", frame.LocalVars)
		fmt.Println("operandstack", frame.OperandStack)
		panic("error")
	}
}

func loop(thread *rtda.Thread, bytecode []byte) {
	frame := thread.PopFrame()
	reader := base.NewCodeReader(bytecode)
	for {
		thread.PC = frame.NextPC
		reader.Reset(thread.PC)
		opcode := reader.ReadUint8()
		inst := instructions.NewInstruction(opcode)
		inst.FetchOperands(reader)
		frame.NextPC = reader.PC()
		fmt.Printf("execute %T %v, pc %d\n", inst, inst, thread.PC)
		inst.Execute(frame)
	}
}
