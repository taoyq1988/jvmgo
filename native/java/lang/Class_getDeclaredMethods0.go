package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

/*
Method(Class<?> declaringClass,
       String name,
       Class<?>[] parameterTypes,
       Class<?> returnType,
       Class<?>[] checkedExceptions,
       int modifiers,
       int slot,
       String signature,
       byte[] annotations,
       byte[] parameterAnnotations,
       byte[] annotationDefault)
*/
const _methodConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"[Ljava/lang/Class;" +
	"Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B[B)V"

func init() {
	_class(getDeclaredMethods0, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;")
}

// private native Method[] getDeclaredMethods0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Method;
func getDeclaredMethods0(frame *rtda.Frame) {
	classObj := frame.GetThis()
	publicOnly := frame.GetBooleanVar(1)

	class := classObj.GetGoClass()
	methods := class.GetMethods(publicOnly)
	methodCount := uint(len(methods))

	methodClass := heap.BootLoader().LoadClass("java/lang/reflect/Method")
	methodArr := methodClass.NewArray(methodCount)

	frame.PushRef(methodArr)

	// create method objs
	if methodCount > 0 {
		thread := frame.Thread
		methodObjs := methodArr.Refs()
		methodConstructor := methodClass.GetConstructor(_methodConstructorDescriptor)
		for i, method := range methods {
			methodObj := methodClass.NewObjWithExtra(method)
			methodObjs[i] = methodObj

			// init methodObj
			thread.InvokeMethodWithShim(methodConstructor, []heap.Slot{
				heap.NewRefSlot(methodObj),                                            // this
				heap.NewRefSlot(classObj),                                             // declaringClass
				heap.NewRefSlot(heap.JSFromGoStr(method.Name)),                        // name
				heap.NewRefSlot(getParameterTypeArr(method)),                          // parameterTypes
				heap.NewRefSlot(getReturnType(method)),                                // returnType
				heap.NewRefSlot(getExceptionTypeArr(method)),                          // checkedExceptions
				heap.NewIntSlot(int32(method.AccessFlag)),                            // modifiers
				heap.NewIntSlot(int32(method.Slot)),                                   // slot
				heap.NewRefSlot(getSignatureStr(method.Signature)),                    // signature
				heap.NewRefSlot(getAnnotationByteArr(method.AnnotationData)),          // annotations
				heap.NewRefSlot(getAnnotationByteArr(method.ParameterAnnotationData)), // parameterAnnotations
				heap.NewRefSlot(getAnnotationByteArr(method.AnnotationDefaultData)),   // annotationDefault
			})
		}
	}
}
