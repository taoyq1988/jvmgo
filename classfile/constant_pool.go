package classfile

import "fmt"

// Constant pool tags
const (
	ConstantUtf8               = 1  // Java 1.0.2
	ConstantInteger            = 3  // Java 1.0.2
	ConstantFloat              = 4  // Java 1.0.2
	ConstantLong               = 5  // Java 1.0.2
	ConstantDouble             = 6  // Java 1.0.2
	ConstantClass              = 7  // Java 1.0.2
	ConstantString             = 8  // Java 1.0.2
	ConstantFieldRef           = 9  // Java 1.0.2
	ConstantMethodRef          = 10 // Java 1.0.2
	ConstantInterfaceMethodRef = 11 // Java 1.0.2
	ConstantNameAndType        = 12 // Java 1.0.2
	ConstantMethodHandle       = 15 // Java 7
	ConstantMethodType         = 16 // Java 7
	ConstantInvokeDynamic      = 18 // Java 7
	ConstantModule             = 19 // Java 9
	ConstantPackage            = 20 // Java 9
	ConstantDynamic            = 17 // Java 11
)

/*
cp_info {
    u1 tag;
    u1 info[];
}
*/
type ConstantInfo interface{}

func readConstantInfo(reader *ClassReader) ConstantInfo {
	tag := reader.ReadUint8()
	switch tag {
	case ConstantInteger:
		return readConstantIntegerInfo(reader)
	case ConstantFloat:
		return readConstantFloatInfo(reader)
	case ConstantLong:
		return readConstantLongInfo(reader)
	case ConstantDouble:
		return readConstantDoubleInfo(reader)
	case ConstantUtf8:
		return readConstantUtf8Info(reader)
	case ConstantString:
		return readConstantStringInfo(reader)
	case ConstantClass:
		return readConstantClassInfo(reader)
	case ConstantFieldRef:
		return readConstantFieldRefInfo(reader)
	case ConstantMethodRef:
		return readConstantMethodRefInfo(reader)
	case ConstantInterfaceMethodRef:
		return readConstantInterfaceMethodRefInfo(reader)
	case ConstantNameAndType:
		return readConstantNameAndTypeInfo(reader)
	case ConstantMethodType:
		return readConstantMethodTypeInfo(reader)
	case ConstantMethodHandle:
		return readConstantMethodHandleInfo(reader)
	case ConstantInvokeDynamic:
		return readConstantInvokeDynamicInfo(reader)
	default:
		panic(fmt.Errorf("invalid constant pool tag: %d", tag))
	}
}

type ConstantPool []ConstantInfo

func parseConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.ReadUint16())
	consts := make([]ConstantInfo, cpCount)

	// The constant_pool table is indexed from 1 to constant_pool_count - 1.
	for i := 1; i < cpCount; i++ {
		consts[i] = readConstantInfo(reader)
		// http://docs.oracle.com/javase/specs/jvms/se8/html/jvms-4.html#jvms-4.4.5
		// All 8-byte constants take up two entries in the constant_pool table of the class file.
		// If a CONSTANT_Long_info or CONSTANT_Double_info structure is the item in the constant_pool
		// table at index n, then the next usable item in the pool is located at index n+2.
		// The constant_pool index n+1 must be valid but is considered unusable.
		switch consts[i].(type) {
		case int64, float64:
			i++
		}
	}

	return ConstantPool{consts}
}
