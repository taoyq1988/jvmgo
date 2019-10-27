package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"github.com/taoyq1988/jvmgo/vmutils"
)

func init() {
	_class(getRawAnnotations, "getRawAnnotations", "()[B")
}

// native byte[] getRawAnnotations();
// ()[B
func getRawAnnotations(frame *rtda.Frame) {
	this := frame.GetThis()

	class := this.GetGoClass()
	goBytes := class.AnnotationData
	if goBytes != nil {
		jBytes := vmutils.CastBytesToInt8s(goBytes)
		byteArr := heap.NewByteArray(jBytes)
		frame.PushRef(byteArr)
		return
	}

	frame.PushRef(nil)
}

// native byte[] getRawTypeAnnotations();
// ()[B
func getRawTypeAnnotations(frame *rtda.Frame) {
	// todo
}
