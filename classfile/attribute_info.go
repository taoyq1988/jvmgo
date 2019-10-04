package classfile

const (
	bootStrapMethods       = "BootstrapMethods"
	code                   = "Code"
	constantValue          = "ConstantValue"
	deprecated             = "Deprecated"
	enclosingMethod        = "EnclosingMethod"
	exceptions             = "Exceptions"
	innerClasses           = "InnerClasses"
	lineNumberTable        = "LineNumberTable"
	localVariableTable     = "LocalVariableTable"
	localVariableTypeTable = "LocalVariableTypeTable"
	signature              = "Signature"
	sourceFile             = "SourceFile"
	synthetic              = "Synthetic"
)

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface{}

type UnparsedAttribute struct {
	Name   string
	Length uint32
	Info   []byte
}

func readAttributes(reader *ClassReader) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttributeInfo(reader)
	}
	return attributes
}

func readAttributeInfo(reader *ClassReader) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrLen := reader.readUint32()
	attrName := reader.cp.getUtf8(attrNameIndex)

	switch attrName {
	case bootStrapMethods:
		return readBootstrapMethodsAttribute(reader)
	case code:
		return readCodeAttribute(reader)
	case constantValue:
		return readConstantValueAttribute(reader)
	case deprecated:
		return DeprecatedAttribute{}
	case enclosingMethod:
		return readEnclosingMethodAttribute(reader)
	case exceptions:
		return readExceptionsAttribute(reader)
	case innerClasses:
		return readInnerClassesAttribute(reader)
	case lineNumberTable:
		return readLineNumberTableAttribute(reader)
	case localVariableTable:
		return readLocalVariableTableAttribute(reader)
	case localVariableTypeTable:
		return readLocalVariableTypeTableAttribute(reader)
	case signature:
		return readSignatureAttribute(reader)
	case sourceFile:
		return readSourceFileAttribute(reader)
	case synthetic:
		return SyntheticAttribute{}
	default:
		// undefined attr
		return UnparsedAttribute{
			Name:   attrName,
			Length: attrLen,
			Info:   reader.readBytes(attrLen),
		}
	}
}
