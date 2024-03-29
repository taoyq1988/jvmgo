package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

func init() {
	_thread(currentThread, "currentThread", "()Ljava/lang/Thread;")
	//_thread(sleep, "sleep", "(J)V")
}

func _thread(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Thread", name, desc, method)
}

// public static native boolean holdsLock(Object obj);
// public static native void yield();
// private native static StackTraceElement[][] dumpThreads(Thread[] threads);
// private native static Thread[] getThreads();

// public static native Thread currentThread();
// ()Ljava/lang/Thread;
func currentThread(frame *rtda.Frame) {
	jThread := frame.Thread.JThread
	frame.PushRef(jThread)
}

// public static native void sleep(long millis) throws InterruptedException;
// (J)V
func sleep(frame *rtda.Frame) {
	//millis := frame.GetLongVar(0)
	//
	//thread := frame.Thread
	//if millis < 0 {
	//	//thread.ThrowIllegalArgumentException("timeout value is negative")
	//	return
	//}
	//
	//m := millis * int64(time.Millisecond)
	//d := time.Duration(m)
	//interrupted := thread.Sleep(d)
	//
	//if interrupted {
	//	//thread.ThrowInterruptedException("sleep interrupted")
	//}
}
