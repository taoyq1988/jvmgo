package misc

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

const (
	vm = "sun/misc/VM"
)

func init() {
	_vm(initialize, "initialize", "()V")
}

func _vm(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(vm, name, desc, method)
}

// private static native void initialize();
// ()V
func initialize(frame *rtda.Frame) {

}
