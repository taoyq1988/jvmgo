package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

const (
	class = "java/lang/Class"
)

func init() {
	//_class(isInstance, "isInstance", "(Ljava/lang/Object;)Z")
	//_class(isInterface, "isInterface", "()Z")
	_class(isAssignableFrom, "isAssignableFrom", "(Ljava/lang/Class;)Z")
	_class(isPrimitive, "isPrimitive", "()Z")
	_class(desiredAssertionStatus0, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z")
	_class(getPrimitiveClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;")
}

func _class(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod(class, name, desc, method)
}

// private static native boolean desiredAssertionStatus0(Class<?> clazz);
// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
	// todo
	//frame.PopRef() // this
	frame.PushBoolean(false)
}

// static native Class<?> getPrimitiveClass(String name);
// (Ljava/lang/String;)Ljava/lang/Class;
func getPrimitiveClass(frame *rtda.Frame) {
	newObj := frame.GetRefVar(0)
	name := heap.JSToGoStr(newObj)
	classLoader := frame.GetClassLoader()
	class := classLoader.GetPrimitiveClass(name)
	classObj := class.JClass
	frame.PushRef(classObj)
}

// public native boolean isAssignableFrom(Class<?> cls);
// (Ljava/lang/Class;)Z
func isAssignableFrom(frame *rtda.Frame) {
	this := frame.GetThis()
	cls := frame.GetRefVar(1)

	thisClass := this.GetGoClass()
	clsClass := cls.GetGoClass()
	ok := thisClass.IsAssignableFrom(clsClass)

	frame.PushBoolean(ok)
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	class := _popClass(frame)
	frame.PushBoolean(class.IsPrimitive())
}

func _popClass(frame *rtda.Frame) *heap.Class {
	this := frame.GetThis()
	return this.GetGoClass()
}
