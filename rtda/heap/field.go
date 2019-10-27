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

// reflection
func (field *Field) Type() *Class {
	if field._type == nil {
		field._type = field.resolveType()
	}
	return field._type
}
func (field *Field) resolveType() *Class {
	descriptor := field.Descriptor
	if len(descriptor) == 1 {
		switch descriptor[0] {
		case 'B':
			return bootLoader.GetPrimitiveClass("byte")
		case 'C':
			return bootLoader.GetPrimitiveClass("char")
		case 'D':
			return bootLoader.GetPrimitiveClass("double")
		case 'F':
			return bootLoader.GetPrimitiveClass("float")
		case 'I':
			return bootLoader.GetPrimitiveClass("int")
		case 'J':
			return bootLoader.GetPrimitiveClass("long")
		case 'S':
			return bootLoader.GetPrimitiveClass("short")
		case 'V':
			return bootLoader.GetPrimitiveClass("void")
		case 'Z':
			return bootLoader.GetPrimitiveClass("boolean")
		default:
			panic("BAD descriptor: " + descriptor)
		}
	}
	if descriptor[0] == 'L' {
		return bootLoader.LoadClass(descriptor[1 : len(descriptor)-1])
	}
	return bootLoader.LoadClass(descriptor)
}
