package heap

type ClassMember struct {
	//AccessFlags
	Name           string
	Descriptor     string
	Signature      string
	AnnotationData []byte // RuntimeVisibleAnnotations_attribute
	Class          *Class
}

type Class struct {
	ConstantPool ConstantPool
}
