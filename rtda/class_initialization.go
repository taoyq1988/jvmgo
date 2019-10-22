package rtda

import (
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"unsafe"
)

//https://docs.oracle.com/javase/specs/jls/se8/html/jls-12.html#jls-12.4.2
func initClass(thread *Thread, class *heap.Class) {
	// step 1
	initCond := class.InitCond
	initCond.L.Lock()

	// step 2 && 3
	threadPtr := uintptr(unsafe.Pointer(thread))
	isInitialized, initThreadPtr := class.IsBeingInitialized()
	if isInitialized {
		if initThreadPtr != threadPtr {
			initCond.Wait()
		} else {
			initCond.L.Unlock()
			return
		}
	}

	// step 4
	if class.IsFullyInitialized() {
		initCond.L.Unlock()
		return
	}

	// step 5
	if class.IsInitializationFailed() {
		initCond.L.Unlock()
		panic("NoClassDefFoundError") //todo throw NoClassDefFoundError
	}

	// step 6
	class.MarkBeingInitialized(threadPtr)
	initCond.L.Unlock()
	initConstantStaticFields(class)

	// step 7
	defer initSuperClass(thread, class)

	// step 8
	//todo

	// step 9 && 10
	callClinit(thread, class)

	//step 11 && 12
	//todo
}

func initSuperClass(thread *Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.SuperClass
		if superClass != nil && superClass.InitializationNotStarted() {
			initClass(thread, superClass)
		}
	}
}

func callClinit(thread *Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit == nil {
		//todo change to do nothing
		clinit = shimReturnMethod
	}

	newFrame := thread.NewFrame(clinit)
	newFrame.AppendOnPopAction(func(popped *Frame) {
		initSucceeded(class)
	})
	thread.PushFrame(newFrame)
}

// step 10
func initSucceeded(class *heap.Class) {
	initCond := class.InitCond
	initCond.L.Lock()
	defer initCond.L.Unlock()

	class.MarkFullyInitialized()
	class.InitCond.Broadcast()
}

func initConstantStaticFields(class *heap.Class) {
	cp := class.ConstantPool
	for _, field := range class.Fields {
		if field.IsStatic() && field.IsFinal() {
			kValIndex := uint(field.ConstValueIndex)
			if kValIndex > 0 {
				slotID := field.SlotID
				staticFields := class.StaticFieldSlots
				switch field.Descriptor {
				case "Z", "B", "C", "S", "I":
					staticFields[slotID] = heap.NewIntSlot(cp.GetConstant(kValIndex).(int32))
				case "J":
					staticFields[slotID] = heap.NewLongSlot(cp.GetConstant(kValIndex).(int64))
				case "F":
					staticFields[slotID] = heap.NewFloatSlot(cp.GetConstant(kValIndex).(float32))
				case "D":
					staticFields[slotID] = heap.NewDoubleSlot(cp.GetConstant(kValIndex).(float64))
				case "Ljava/lang/String;":
					staticFields[slotID] = heap.NewRefSlot(cp.GetConstantString(kValIndex).GetJString())
				}
			}
		}
	}
}
