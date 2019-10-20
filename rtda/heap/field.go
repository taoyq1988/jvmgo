package heap

import "github.com/taoyq1988/jvmgo/classfile"

type Field struct {
	ClassMember
	IsLongOrDouble  bool
	ConstValueIndex uint16
	SlotID          uint
	_type           *Class
}

func newField(class *Class, cf *classfile.Classfile, cfMember classfile.MemberInfo) *Field {
	field := &Field{}
	field.Class = class
	field.parseMemberData(cf, cfMember)
	field.IsLongOrDouble = field.Descriptor == "J" || field.Descriptor == "D"
	field.ConstValueIndex = cfMember.GetConstantValueIndex()
	return field
}
