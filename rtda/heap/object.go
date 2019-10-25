package heap

import (
	"reflect"
	"sync"
)

type Object struct {
	Class  *Class
	Fields interface{} // []Slot for Object, []int32 for int[] ...
	Extra  interface{} // remember some important things from Golang
	//Fixme: can simple
	Monitor *Monitor
	lock    *sync.RWMutex // state lock
}

func newObj(class *Class, fields, extra interface{}) *Object {
	return &Object{
		Class:   class,
		Fields:  fields,
		Extra:   extra,
		Monitor: newMonitor(),
		lock:    &sync.RWMutex{},
	}
}

func (object *Object) initFields() {
	fields := object.Fields.([]Slot)
	for class := object.Class; class != nil; class = class.SuperClass {
		for _, f := range class.Fields {
			if !f.IsStatic() {
				fields[f.SlotID] = EmptySlot
			}
		}
	}
}

/**
Object Array
*/
func NewArray(arrClass *Class, count uint) *Object {
	if arrClass.IsPrimitiveArray() {
		return newPrimitiveArray(arrClass, count)
	}
	componentClass := arrClass.ComponentClass()
	return NewRefArray(componentClass, count)
}

func NewPrimitiveArray(aType uint8, count uint) *Object {
	class := getPrimitiveClassByType(aType)
	return newPrimitiveArray(class, count)
}

func newPrimitiveArray(arrClass *Class, count uint) *Object {
	switch arrClass.Name {
	case BooleanPrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]int8, count), nil)
	case BytePrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]int8, count), nil)
	case CharPrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]uint16, count), nil)
	case ShortPrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]int16, count), nil)
	case IntPrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]int32, count), nil)
	case LongPrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]int64, count), nil)
	case FloatPrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]float32, count), nil)
	case DoublePrimitiveType.ArrayClassName:
		return newObj(arrClass, make([]float64, count), nil)
	default:
		panic("not primitive array " + arrClass.Name)
	}
}

func NewByteArray(bytes []int8) *Object {
	return newObj(bootLoader.getClass("[B"), bytes, nil)
}

func NewCharArray(chars []uint16) *Object {
	return newObj(bootLoader.getClass("[C"), chars, nil)
}

func NewRefArray(componentClass *Class, count uint) *Object {
	arrClass := componentClass.arrayClass()
	components := make([]*Object, count)
	return newObj(arrClass, components, nil)
}

func ArrayLength(arr *Object) int32 {
	return int32(reflect.ValueOf(arr.Fields).Len())
}

func (obj *Object) Refs() []*Object    { return obj.Fields.([]*Object) }
func (obj *Object) Booleans() []int8   { return obj.Fields.([]int8) }
func (obj *Object) Bytes() []int8      { return obj.Fields.([]int8) }
func (obj *Object) Chars() []uint16    { return obj.Fields.([]uint16) }
func (obj *Object) Shorts() []int16    { return obj.Fields.([]int16) }
func (obj *Object) Ints() []int32      { return obj.Fields.([]int32) }
func (obj *Object) Longs() []int64     { return obj.Fields.([]int64) }
func (obj *Object) Floats() []float32  { return obj.Fields.([]float32) }
func (obj *Object) Doubles() []float64 { return obj.Fields.([]float64) }

/**
InstanceOf
*/
func (obj *Object) IsInstanceOf(class *Class) bool {
	s, t := obj.Class, class
	return checkcast(s, t)
}

func checkcast(s, t *Class) bool {
	if s == t {
		return true
	}

	if s.IsArray() {
		if t.IsArray() {
			sc := s.ComponentClass()
			tc := t.ComponentClass()
			return sc == tc || checkcast(sc, tc)
		} else {
			if t.IsInterface() {
				return t.isJlCloneable() || t.isJioSerializable()
			} else {
				return t.isJlObject()
			}
		}
	} else {
		if s.IsInterface() {
			if t.IsInterface() {
				return t.isSuperInterfaceOf(s)
			} else {
				return t.isJlObject()
			}
		} else {
			if t.IsInterface() {
				return s.IsImplements(t)
			} else {
				return s.isSubClassOf(t)
			}
		}
	}
}

/**
reflection
*/
func (obj *Object) GetFieldValue(fieldName, descriptor string) Slot {
	field := obj.Class.GetInstanceField(fieldName, descriptor)
	return field.GetValue(obj)
}

func (obj *Object) SetFieldValue(fieldName, descriptor string, value Slot) {
	field := obj.Class.GetInstanceField(fieldName, descriptor)
	field.PutValue(obj, value)
}
