package heap

import (
	"github.com/taoyq1988/jvmgo/classfile"
	"strings"
)

func getNameAndType(cf *classfile.Classfile, index uint16) (name, _type string) {
	if index > 0 {
		ntInfo := cf.GetConstantInfo(index).(classfile.ConstantNameAndTypeInfo)
		name = cf.GetUTF8(ntInfo.NameIndex)
		_type = cf.GetUTF8(ntInfo.DescriptorIndex)
	}
	return
}

func DotToSlash(name string) string {
	return strings.ReplaceAll(name, ".", "/")
}
func SlashToDot(name string) string {
	return strings.ReplaceAll(name, "/", ".")
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	if className[0] == '[' {
		// array
		return "[" + className
	}
	for _, primitiveType := range PrimitiveTypes {
		if primitiveType.Name == className {
			// primitive
			return primitiveType.ArrayClassName
		}
	}
	// object
	return "[L" + className + ";"
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	if className[0] == '[' {
		descriptor := className[1:]
		return getClassName(descriptor)
	}
	panic("Not array: " + className)
}

// [XXX -> [XXX
// LXXX; -> XXX
// I -> int ...
func getClassName(descriptor string) string {
	switch descriptor[0] {
	case '[': // array
		return descriptor
	case 'L': // object
		return descriptor[1 : len(descriptor)-1]
	default: // primitive
		return getPrimitiveType(descriptor)
	}
}
