package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
	"strings"
)

func init() {
	//_class(getClassLoader0, "getClassLoader0", "()Ljava/lang/ClassLoader;")
	_class(getComponentType, "getComponentType", "()Ljava/lang/Class;")
	_class(getConstantPool, "getConstantPool", "()Lsun/reflect/ConstantPool;")
	_class(getDeclaringClass0, "getDeclaringClass0", "()Ljava/lang/Class;")
	//_class(getEnclosingMethod0, "getEnclosingMethod0", "()[Ljava/lang/Object;")
	//_class(getInterfaces0, "getInterfaces0", "()[Ljava/lang/Class;")
	_class(getModifiers, "getModifiers", "()I")
	_class(getName0, "getName0", "()Ljava/lang/String;")
	_class(getSuperclass, "getSuperclass", "()Ljava/lang/Class;")
	_class(isArray, "isArray", "()Z")
	_class(isAssignableFrom, "isAssignableFrom", "(Ljava/lang/Class;)Z")
	_class(isInstance, "isInstance", "(Ljava/lang/Object;)Z")
	_class(isInterface, "isInterface", "()Z")
	_class(isPrimitive, "isPrimitive", "()Z")
	_class(getGenericSignature0, "getGenericSignature0", "()Ljava/lang/String;")
}

func _class(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("java/lang/Class", name, desc, method)
}

// native ClassLoader getClassLoader0();
// ()Ljava/lang/ClassLoader;
func getClassLoader0(frame *rtda.Frame) {
	//class := _popClass(frame)
	//from := class.LoadedFrom
	//
	//if cp.IsBootClassPath(from, frame.Thread.VMOptions.AbsJreLib) {
	//	frame.PushRef(nil)
	//	return
	//}
	//
	//clClass := heap.BootLoader().LoadClass("java/lang/ClassLoader")
	//getSysCl := clClass.GetStaticMethod("getSystemClassLoader", "()Ljava/lang/ClassLoader;")
	//frame.Thread.InvokeMethod(getSysCl)
}

// public native Class<?> getComponentType();
// ()Ljava/lang/Class;
func getComponentType(frame *rtda.Frame) {
	class := _popClass(frame)
	componentClass := class.ComponentClass()
	componentClassObj := componentClass.JClass

	frame.PushRef(componentClassObj)
}

// native ConstantPool getConstantPool();
// ()Lsun/reflect/ConstantPool;
func getConstantPool(frame *rtda.Frame) {
	class := _popClass(frame)
	cpClass := heap.BootLoader().LoadClass("sun/reflect/ConstantPool")
	if cpClass.InitializationNotStarted() {
		frame.RevertNextPC()
		frame.Thread.InitClass(cpClass)
		return
	}

	cpObj := cpClass.NewObjWithExtra(class.ConstantPool) // todo init cpObj
	frame.PushRef(cpObj)
}

// private native Class<?> getDeclaringClass0();
// ()Ljava/lang/Class;
func getDeclaringClass0(frame *rtda.Frame) {
	class := _popClass(frame)
	if class.IsArray() || class.IsPrimitive() {
		frame.PushRef(nil)
		return
	}

	lastDollarIndex := strings.LastIndex(class.Name, "$")
	if lastDollarIndex < 0 {
		frame.PushRef(nil)
		return
	}

	// todo
	declaringClassName := class.Name[:lastDollarIndex]
	declaringClass := frame.GetClassLoader().LoadClass(declaringClassName)
	frame.PushRef(declaringClass.JClass)
}

// private native Object[] getEnclosingMethod0();
// ()[Ljava/lang/Object;
func getEnclosingMethod0(frame *rtda.Frame) {
	class := _popClass(frame)
	if class.IsPrimitive() {
		frame.PushNull()
	} else {
		emInfo := class.EnclosingMethod
		emInfoObj := _createEnclosintMethodInfo(frame.GetClassLoader(), emInfo)
		if emInfoObj == nil || heap.ArrayLength(emInfoObj) == 0 {
			frame.PushNull()
		} else {
			frame.PushRef(emInfoObj)
		}
	}
}

func _createEnclosintMethodInfo(classLoader *heap.ClassLoader, emInfo *heap.EnclosingMethod) *heap.Object {
	//if emInfo == nil {
	//	return nil
	//}
	//
	//enclosingClass := classLoader.LoadClass(emInfo.ClassName)
	//enclosingClassObj := enclosingClass.JClass
	//var methodNameObj, methodDescriptorObj *heap.Object
	//if emInfo.MethodName != "" {
	//	methodNameObj = heap.JSFromGoStr(emInfo.MethodName)
	//	methodDescriptorObj = heap.JSFromGoStr(emInfo.MethodDescriptor)
	//} else {
	//	methodNameObj, methodDescriptorObj = nil, nil
	//}
	//
	//objs := []*heap.Object{enclosingClassObj, methodNameObj, methodDescriptorObj}
	//return heap.NewRefArray2(classLoader.JLObjectClass(), objs) // Object[]
	return nil
}

// private native Class<?>[] getInterfaces0();
// ()[Ljava/lang/Class;
func getInterfaces0(frame *rtda.Frame) {
	//class := _popClass(frame)
	//interfaces := class.Interfaces
	//interfaceObjs := make([]*heap.Object, len(interfaces))
	//for i, iface := range interfaces {
	//	interfaceObjs[i] = iface.JClass
	//}
	//
	//jlClassClass := heap.BootLoader().JLClassClass()
	//interfaceArr := heap.NewRefArray2(jlClassClass, interfaceObjs)
	//
	//frame.PushRef(interfaceArr)
}

// private native String getName0();
// ()Ljava/lang/String;
func getName0(frame *rtda.Frame) {
	class := _popClass(frame)
	name := class.NameJlsFormat()
	nameObj := heap.JSFromGoStr(name)

	frame.PushRef(nameObj)
}

// public native int getModifiers();
// ()I
func getModifiers(frame *rtda.Frame) {
	class := _popClass(frame)
	frame.PushInt(int32(class.AccessFlag))
}

// public native Class<? super T> getSuperclass();
// ()Ljava/lang/Class;
func getSuperclass(frame *rtda.Frame) {
	class := _popClass(frame)
	superClass := class.SuperClass

	if superClass != nil {
		frame.PushRef(superClass.JClass)
	} else {
		frame.PushNull()
	}
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

// public native boolean isInstance(Object obj);
// (Ljava/lang/Object;)Z
func isInstance(frame *rtda.Frame) {
	this := frame.GetThis()
	obj := frame.GetRefVar(1)

	class := this.GetGoClass()
	ok := obj.IsInstanceOf(class)

	frame.PushBoolean(ok)
}

// public native boolean isArray();
// ()Z
func isArray(frame *rtda.Frame) {
	class := _popClass(frame)
	frame.PushBoolean(class.IsArray())
}

// public native boolean isInterface();
// ()Z
func isInterface(frame *rtda.Frame) {
	class := _popClass(frame)
	frame.PushBoolean(class.IsInterface())
}

// public native boolean isPrimitive();
// ()Z
func isPrimitive(frame *rtda.Frame) {
	class := _popClass(frame)
	frame.PushBoolean(class.IsPrimitive())
}

// private native String getGenericSignature0();
// ()Ljava/lang/String;
func getGenericSignature0(frame *rtda.Frame) {
	class := _popClass(frame)
	if class == nil {
		panic("illegal class")
	}

	// Return null for arrays and primatives
	if !class.IsPrimitive() {
		signature := class.Signature
		if signature == "" {
			frame.PushNull()
		} else {
			frame.PushRef(heap.JSFromGoStr(signature))
		}
		return
	}

	frame.PushNull()
}

func _popClass(frame *rtda.Frame) *heap.Class {
	this := frame.GetThis()
	return this.GetGoClass()
}
