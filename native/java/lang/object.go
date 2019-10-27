package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"unsafe"
)

func init() {
	_object(clone, "clone", "()Ljava/lang/Object;")
	_object(getClass, "getClass", "()Ljava/lang/Class;")
	_object(hashCode, "hashCode", "()I")
	_object(notifyAll, "notifyAll", "()V")
	_object(wait, "wait", "(J)V")
}

func _object(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Object", name, desc, method)
}

// protected native Object clone() throws CloneNotSupportedException;
// ()Ljava/lang/Object;
func clone(frame *rtda.Frame) {
	this := frame.GetThis()

	//todo
	frame.PushRef(this.Clone())
}

// public final native Class<?> getClass();
// ()Ljava/lang/Class;
func getClass(frame *rtda.Frame) {
	this := frame.GetThis()

	class := this.Class.JClass
	frame.PushRef(class)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
	this := frame.GetThis()

	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.PushInt(hash)
}

// public final native void notify();

// public final native void notifyAll();
// ()V
func notifyAll(frame *rtda.Frame) {
	this := frame.GetThis()

	thread := frame.Thread
	monitor := this.Monitor
	if !monitor.HasOwner(thread) {
		// todo
		panic("IllegalMonitorStateException")
	}

	monitor.NotifyAll()
}

// public final native void wait(long timeout) throws InterruptedException;
// (J)V
func wait(frame *rtda.Frame) {
	this := frame.GetThis()
	// timeout := frame.GetLongVar(1) // todo

	thread := frame.Thread
	monitor := this.Monitor
	if !monitor.HasOwner(thread) {
		// todo
		panic("IllegalMonitorStateException")
	}

	monitor.Wait()
}
