package interpret

import (
	"fmt"
	"github.com/taoyq1988/jvmgo/instructions"
	"github.com/taoyq1988/jvmgo/instructions/base"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

const isLog = false

var _mainThreadGroup *heap.Object
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
	initBootClass(thread)
	thread.PushFrame(frame)
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

func mainThreadNotReady(thread *rtda.Thread) bool {
	classLoader := heap.BootLoader()
	frame := thread.CurrentFrame()
	if _mainThreadGroup == nil {
		undoExec(thread)
		threadGroupClass := classLoader.LoadClass("java/lang/ThreadGroup")
		_mainThreadGroup = threadGroupClass.NewObj()
		initMethod := threadGroupClass.GetConstructor("()V")
		frame.PushRef(_mainThreadGroup) // this
		thread.InvokeMethod(initMethod)
		return true
	}
	if thread.JThread == nil {
		undoExec(thread)
		threadClass := classLoader.LoadClass("java/lang/Thread")
		mainThreadObj := threadClass.NewObjWithExtra(thread)
		mainThreadObj.SetFieldValue("priority", "I", heap.NewIntSlot(1))
		thread.JThread = mainThreadObj

		initMethod := threadClass.GetConstructor("(Ljava/lang/ThreadGroup;Ljava/lang/String;)V")
		frame.PushRef(mainThreadObj)            // this
		frame.PushRef(_mainThreadGroup)         // group
		frame.PushRef(heap.JSFromGoStr("main")) // name
		thread.InvokeMethod(initMethod)
		return true
	}
	return false
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
	fmt.Printf("[execute] method %s, method desc %s, class %s, inst %T %v, pc %d\n", frame.Method.Name, frame.Method.Descriptor, frame.Method.Class.Name, inst, inst, frame.Thread.PC)
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
