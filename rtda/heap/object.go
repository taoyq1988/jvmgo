package heap

type Object struct {
	Class   *Class
	Fields  interface{} // []Slot for Object, []int32 for int[] ...
	Extra   interface{} // remember some important things from Golang
}
