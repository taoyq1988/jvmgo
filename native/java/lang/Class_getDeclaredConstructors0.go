package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

/*
Constructor(Class<T> declaringClass,
            Class<?>[] parameterTypes,
            Class<?>[] checkedExceptions,
            int modifiers,
            int slot,
            String signature,
            byte[] annotations,
            byte[] parameterAnnotations)
}
*/
const _constructorConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B)V"

func init() {
	_class(getDeclaredConstructors0, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;")
}

// private native Constructor<T>[] getDeclaredConstructors0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Constructor;
func getDeclaredConstructors0(frame *rtda.Frame) {
	classObj := frame.GetThis()
	publicOnly := frame.GetBooleanVar(1)

	class := classObj.GetGoClass()
	constructors := class.GetConstructors(publicOnly)
	constructorCount := uint(len(constructors))

	constructorClass := heap.BootLoader().LoadClass("java/lang/reflect/Constructor")
	constructorArr := constructorClass.NewArray(constructorCount)

	frame.PushRef(constructorArr)

	if constructorCount > 0 {
		thread := frame.Thread
		constructorObjs := constructorArr.Refs()
		constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
		for i, constructor := range constructors {
			constructorObj := constructorClass.NewObjWithExtra(constructor)
			constructorObjs[i] = constructorObj

			// init constructorObj
			thread.InvokeMethodWithShim(constructorInitMethod, []heap.Slot{
				heap.NewRefSlot(constructorObj),                                            // this
				heap.NewRefSlot(classObj),                                                  // declaringClass
				heap.NewRefSlot(getParameterTypeArr(constructor)),                          // parameterTypes
				heap.NewRefSlot(getExceptionTypeArr(constructor)),                          // checkedExceptions
				heap.NewIntSlot(int32(constructor.AccessFlag)),                            // modifiers
				heap.NewIntSlot(int32(0)),                                                  // todo slot
				heap.NewRefSlot(getSignatureStr(constructor.Signature)),                    // signature
				heap.NewRefSlot(getAnnotationByteArr(constructor.AnnotationData)),          // annotations
				heap.NewRefSlot(getAnnotationByteArr(constructor.ParameterAnnotationData)), // parameterAnnotations
			})
		}
	}
}
