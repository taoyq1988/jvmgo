package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"strings"

)

func init() {
	_class(desiredAssertionStatus0, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z")
	_class(forName0, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;")
	_class(getPrimitiveClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;")
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	//frame.PopRef() // this
	frame.PushBoolean(false)
}

// private static native Class<?> forName0(String name, boolean initialize,
//                                         ClassLoader loader,
//                                         Class<?> caller)
//     throws ClassNotFoundException;
// (Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;
func forName0(frame *rtda.Frame) {
	jName := frame.GetRefVar(0)
	initialize := frame.GetBooleanVar(1)
	//jLoader := frame.GetRefVar(2)

	goName := heap.JSToGoStr(jName)
	goName = strings.ReplaceAll(goName, ".", "/")
	goClass := frame.GetClassLoader().LoadClass(goName)
	jClass := goClass.JClass

	if initialize && goClass.InitializationNotStarted() {
		// undo forName0
		thread := frame.Thread
		frame.NextPC = thread.PC
		// init class
		thread.InitClass(goClass)
	} else {
		frame.PushRef(jClass)
	}
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
	nameObj := frame.GetRefVar(0)

	name := heap.JSToGoStr(nameObj)
	classLoader := frame.GetClassLoader()
	class := classLoader.GetPrimitiveClass(name)
	classObj := class.JClass

	frame.PushRef(classObj)
}
