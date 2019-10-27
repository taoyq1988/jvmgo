package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"math"
)

const (
	float = "java/lang/Float"
)

func init() {
	_float(floatToRawIntBits, "floatToRawIntBits", "(F)I")
}

func _float(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(float, name, desc, method)
}

// public static native int floatToRawIntBits(float value);
// (F)I
func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.GetFloatVar(0)
	bits := math.Float32bits(value)

	frame.PushInt(int32(bits)) // todo
}
