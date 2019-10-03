package classfile

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
	return nil
}
