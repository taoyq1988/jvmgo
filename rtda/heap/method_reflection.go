package heap

func (method *Method) GetParameterTypes() []*Class {
	if method.ParamSlotCount == 0 {
		return nil
	}

	paramClasses := make([]*Class, 0, method.ParamSlotCount)
	for _, paramType := range method.ParameterTypes {
		paramClassName := getClassName(string(paramType))
		paramClasses = append(paramClasses, bootLoader.LoadClass(paramClassName))
	}

	return paramClasses
}

func (method *Method) GetReturnType() *Class {
	returnDescriptor := method.ReturnType
	returnClassName := getClassName(string(returnDescriptor))
	returnClass := bootLoader.LoadClass(returnClassName)
	return returnClass
}

func (method *Method) GetExceptionTypes() []*Class {
	if method.exceptionTableIndex == nil {
		return nil
	}

	exClasses := make([]*Class, len(method.exceptionTableIndex))
	cp := method.Class.ConstantPool

	for i, exIndex := range method.exceptionTableIndex {
		kClass := cp.GetConstantClass(uint(exIndex))
		exClasses[i] = kClass.GetClass()
	}

	return exClasses
}
