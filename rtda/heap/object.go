package heap

import "sync"

type Object struct {
	Class  *Class
	Fields interface{} // []Slot for Object, []int32 for int[] ...
	Extra  interface{} // remember some important things from Golang
	//Fixme: can simple
	Monitor *Monitor
	lock    *sync.RWMutex // state lock
}

func newObj(class *Class, fields, extra interface{}) *Object {
	return &Object{class, fields, extra, newMonitor(), &sync.RWMutex{}}
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
