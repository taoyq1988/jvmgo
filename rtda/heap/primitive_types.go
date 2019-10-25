package heap

import "fmt"

const (
	//Array Type  atype
	ATBoolean = 4
	ATChar    = 5
	ATFloat   = 6
	ATDouble  = 7
	ATByte    = 8
	ATShort   = 9
	ATInt     = 10
	ATLong    = 11
)

var (
	VoidPrimitiveType = PrimitiveType{"V", "[V", "void", "java/lang/Void"}
	BooleanPrimitiveType = PrimitiveType{"Z", "[Z", "boolean", "java/lang/Boolean"}
	BytePrimitiveType = PrimitiveType{"B", "[B", "byte", "java/lang/Byte"}
	CharPrimitiveType = PrimitiveType{"C", "[C", "char", "java/lang/Character"}
	ShortPrimitiveType = PrimitiveType{"S", "[S", "short", "java/lang/Short"}
	IntPrimitiveType = PrimitiveType{"I", "[I", "int", "java/lang/Integer"}
	LongPrimitiveType = PrimitiveType{"J", "[J", "long", "java/lang/Long"}
	FloatPrimitiveType = PrimitiveType{"F", "[F", "float", "java/lang/Float"}
	DoublePrimitiveType = PrimitiveType{"D", "[D", "double", "java/lang/Double"}

	PrimitiveTypes = []PrimitiveType{
		VoidPrimitiveType,
		BooleanPrimitiveType,
		BytePrimitiveType,
		CharPrimitiveType,
		ShortPrimitiveType,
		IntPrimitiveType,
		LongPrimitiveType,
		FloatPrimitiveType,
		DoublePrimitiveType,
	}
)

// type jboolean bool
// type jbyte int8
// type jchar uint16
// type jshort int16
// type jint int32
// type jlong int64
// type jfloat float32
// type jdouble float64

type PrimitiveType struct {
	Descriptor       string
	ArrayClassName   string
	Name             string
	WrapperClassName string
}

func isPrimitiveType(name string) bool {
	for _, primitiveType := range PrimitiveTypes {
		if primitiveType.Name == name {
			return true
		}
	}
	return false
}

func getPrimitiveType(descriptor string) string {
	for _, primitiveType := range PrimitiveTypes {
		if primitiveType.Descriptor == descriptor {
			return primitiveType.Name
		}
	}
	panic("Not primitive type: " + descriptor)
}

func getPrimitiveClassByType(aType uint8) *Class {
	switch aType {
	case ATBoolean:
		return bootLoader.getClass(VoidPrimitiveType.ArrayClassName)
	case ATByte:
		return bootLoader.getClass(BooleanPrimitiveType.ArrayClassName)
	case ATChar:
		return bootLoader.getClass(CharPrimitiveType.ArrayClassName)
	case ATShort:
		return bootLoader.getClass(ShortPrimitiveType.ArrayClassName)
	case ATInt:
		return bootLoader.getClass(IntPrimitiveType.ArrayClassName)
	case ATLong:
		return bootLoader.getClass(LongPrimitiveType.ArrayClassName)
	case ATFloat:
		return bootLoader.getClass(FloatPrimitiveType.ArrayClassName)
	case ATDouble:
		return bootLoader.getClass(DoublePrimitiveType.ArrayClassName)
	default:
		panic(fmt.Errorf("invalid atype: %v", aType))
	}
}
