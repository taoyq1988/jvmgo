package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

const (
	system = "java/lang/System"
)

func init() {
	_system(initProperties, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;")
	_system(arraycopy, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V")
}

func _system(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(system, name, desc, method)
}

// private static native Properties initProperties(Properties props);
// (Ljava/util/Properties;)Ljava/util/Properties;
func initProperties(frame *rtda.Frame) {
	props := frame.GetRefVar(0)

	frame.PushRef(props)
}

func arraycopy(frame *rtda.Frame) {

}
