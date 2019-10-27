package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"unsafe"
)

const(
	object = "java/lang/Object"
)

func init() {
	_object(hashCode, "hashCode", "()I")
}

func _object(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(object, name, desc, method)
}

// public native int hashCode();
// ()I
func hashCode(frame *rtda.Frame) {
	this := frame.GetThis()

	hash := int32(uintptr(unsafe.Pointer(this)))
	frame.PushInt(hash)
}
