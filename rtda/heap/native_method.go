package heap

var (
	nativeRegistry    = map[string]interface{}{}
	emptyNativeMethod interface{}
)

func SetEmptyNativeMethod(m interface{}) {
	emptyNativeMethod = m
}

func RegisterNativeMethod(className, methodName, descriptor string, method interface{}) {
	key := className + "@" + methodName + "@" + descriptor
	if _, ok := nativeRegistry[key]; !ok {
		nativeRegistry[key] = method
	}
}

func findNativeMethod(method *Method) interface{} {
	key := method.Class.Name + "@" + method.Name + "@" + method.Descriptor
	if method, ok := nativeRegistry[key]; ok {
		return method
	}
	if method.IsRegisterNatives() || method.IsInitIDs() {
		return emptyNativeMethod
	}
	panic("can not find native method " + key)
}
