package interpret

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/instructions"
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

const isLog = true

var _bootClasses = []string{
	"java/lang/Class",
	"java/lang/String",
	"java/lang/System",
	"java/lang/Thread",
	"java/lang/ThreadGroup",
	"java/io/PrintStream",
}

func Interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	initBootClass(thread)
	jlSystemNotReady(thread)
	defer catchError(frame)
	loop(thread)
}

func initBootClass(thread *rtda.Thread) {
	classLoader := heap.BootLoader()
	for _, className := range _bootClasses {
		class := classLoader.LoadClass(className)
		if class.InitializationNotStarted() {
			thread.InitClass(class)
		}
	}
}

func jlSystemNotReady(thread *rtda.Thread) bool {
	classLoader := heap.BootLoader()
	sysClass := classLoader.LoadClass("java/lang/System")
	propsField := sysClass.GetStaticField("props", "Ljava/util/Properties;")
	props := propsField.GetStaticValue().Ref
	if props == nil {
		undoExec(thread)
		initSys := sysClass.GetStaticMethod("initializeSystemClass", "()V")
		thread.InvokeMethod(initSys)
		return true
	}
	return false
}

func undoExec(thread *rtda.Thread) {
	thread.CurrentFrame().RevertNextPC()
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
