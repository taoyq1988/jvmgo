package heap

import "github.com/taoyq1988/jvmgo/classfile"

func newClass(cf *classfile.Classfile) *Class {
	class := &Class{}
	class.AccessFlag = AccessFlag(cf.AccessFlags)

	class.parseClassNames(cf)
	class.parseFields(cf)

	return nil
}

func (class *Class) parseConstantPool(cf *classfile.Classfile) {

}

func (class *Class) parseClassNames(cf *classfile.Classfile) {
	class.Name = cf.GetThisClassName()
	class.superClassName = cf.GetSuperClassName()
	class.interfaceNames = cf.GetInterfaceNames()
}

func (class *Class) parseFields(cf *classfile.Classfile) {
	class.Fields = make([]*Field, len(cf.Fields))
	for i, fieldInfo := range cf.Fields {
		class.Fields[i] = newField(class, cf, fieldInfo)
	}
}

func (class *Class) parseMethods(cf *classfile.Classfile) {

}

func (class *Class) parseAttributes(cf *classfile.Classfile) {

}
