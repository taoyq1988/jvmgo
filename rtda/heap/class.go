package heap

import (
	"github.com/taoyq1988/jvmgo/classfile"
	"github.com/taoyq1988/jvmgo/classpath"
)

type ClassMember struct {
	AccessFlag
	Name           string
	Descriptor     string
	Signature      string
	AnnotationData []byte // RuntimeVisibleAnnotations_attribute
	Class          *Class
}

func (member *ClassMember) parseMemberData(cf *classfile.Classfile, cfMember classfile.MemberInfo) {
	member.AccessFlag = AccessFlag(cfMember.AccessFlags)
	member.Name = cf.GetUTF8(cfMember.NameIndex)
	member.Descriptor = cf.GetUTF8(cfMember.DescriptorIndex)
	member.Signature = cf.GetUTF8(cfMember.GetSignatureIndex())
	member.AnnotationData = cfMember.GetRuntimeVisibleAnnotationsAttributeData()
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

func (class *Class) String() string {
	return `{Class name:` + class.Name + `}`
}
