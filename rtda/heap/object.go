package heap

type Object struct {
	Class   *Class
	Fields  interface{} // []Slot for Object, []int32 for int[] ...
	Extra   interface{} // remember some important things from Golang
}

func newObj(class *Class, fields, extra interface{}) *Object {
	return &Object{class, fields, extra}
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
