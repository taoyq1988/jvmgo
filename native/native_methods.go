package native

import (
	_ "github.com/taoyq1988/jvmgo/native/java/io"
	_ "github.com/taoyq1988/jvmgo/native/java/lang"
	_ "github.com/taoyq1988/jvmgo/native/sun/misc"
	_ "github.com/taoyq1988/jvmgo/native/sun/reflect"
	_ "github.com/taoyq1988/jvmgo/native/java/security"
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

func init() {
	heap.SetEmptyNativeMethod(emptyNativeMethod)
}

func emptyNativeMethod(frame *rtda.Frame) {
	// do nothing
}
