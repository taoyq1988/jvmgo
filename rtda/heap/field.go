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

func (field *Field) GetValue(ref *Object) Slot {
	fields := ref.Fields.([]Slot)
	return fields[field.SlotID]
}

func (field *Field) PutValue(ref *Object, val Slot) {
	fields := ref.Fields.([]Slot)
	fields[field.SlotID] = val
}

func (field *Field) GetStaticValue() Slot {
	return field.Class.StaticFieldSlots[field.SlotID]
}

func (field *Field) PutStaticValue(val Slot) {
	field.Class.StaticFieldSlots[field.SlotID] = val
}
