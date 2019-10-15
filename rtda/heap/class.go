package heap

import "github.com/taoyq1988/jvmgo/classpath"

type ClassMember struct {
	//AccessFlags
	Name           string
	Descriptor     string
	Signature      string
	AnnotationData []byte // RuntimeVisibleAnnotations_attribute
	Class          *Class
}

type Class struct {
	AccessFlag
	ConstantPool
	Name               string // thisClassName
	superClassName     string
	interfaceNames     []string
	Fields             []*Field
	Methods            []*Method
	instanceFieldCount uint
	staticFieldCount   uint
	StaticFieldSlots   []Slot
	vtable             []*Method // virtual method table
	JClass             *Object   // java.lang.Class instance
	SuperClass         *Class
	Interfaces         []*Class
	LoadedFrom         classpath.Entry
	initState          int
}
