package lang

import (
	"github.com/taoyq1988/jvmgo/rtda"
	"github.com/taoyq1988/jvmgo/rtda/heap"
)

/*
Field(Class<?> declaringClass,
      String name,
      Class<?> type,
      int modifiers,
      int slot,
      String signature,
      byte[] annotations)
*/
const _fieldConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B)V"

func init() {
	_class(getDeclaredFields0, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;")
}

// private native Field[] getDeclaredFields0(boolean publicOnly);
// (Z)[Ljava/lang/reflect/Field;
func getDeclaredFields0(frame *rtda.Frame) {
	classObj := frame.GetThis()
	publicOnly := frame.GetBooleanVar(1)

	class := classObj.GetGoClass()
	fields := class.GetFields(publicOnly)
	fieldCount := uint(len(fields))

	fieldClass := heap.BootLoader().LoadClass("java/lang/reflect/Field")
	fieldArr := heap.NewRefArray(fieldClass, fieldCount)

	frame.PushRef(fieldArr)

	if fieldCount > 0 {
		thread := frame.Thread
		fieldObjs := fieldArr.Refs()
		fieldConstructor := fieldClass.GetConstructor(_fieldConstructorDescriptor)
		for i, goField := range fields {
			fieldObj := fieldClass.NewObjWithExtra(goField)
			fieldObjs[i] = fieldObj

			// init fieldObj
			thread.InvokeMethodWithShim(fieldConstructor, []heap.Slot{
				heap.NewRefSlot(fieldObj),                                     // this
				heap.NewRefSlot(classObj),                                     // declaringClass
				heap.NewRefSlot(heap.JSFromGoStr(goField.Name)),               // name
				heap.NewRefSlot(goField.Type().JClass),                        // type
				heap.NewIntSlot(int32(goField.AccessFlag)),                   // modifiers
				heap.NewIntSlot(int32(goField.SlotID)),                        // slot
				heap.NewRefSlot(getSignatureStr(goField.Signature)),           // signature
				heap.NewRefSlot(getAnnotationByteArr(goField.AnnotationData)), // annotations
			})
		}
	}
}
