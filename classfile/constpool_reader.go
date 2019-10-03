package classfile

import (
	"fmt"
	"math"
	"unicode/utf16"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantIntegerInfo(reader *ClassReader) int32 {
	return int32(reader.readUint32())
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantFloatInfo(reader *ClassReader) float32 {
	return math.Float32frombits(reader.readUint32())
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantLongInfo(reader *ClassReader) int64 {
	return int64(reader.readUint64())
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantDoubleInfo(reader *ClassReader) float64 {
	return math.Float64frombits(reader.readUint64())
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
func readConstantUtf8Info(reader *ClassReader) string {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	return decodeMUTF8(bytes)
}

/*
func decodeMUTF8(bytes []byte) string {
	return string(bytes) // not correct!
}
*/

// mutf8 -> utf16 -> utf32 -> string
// see java.io.DataInputStream.readUTF(DataInput)
func decodeMUTF8(bytearr []byte) string {
	utflen := len(bytearr)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytearr[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytearr[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			/* 0xxxxxxx*/
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			/* 110x xxxx   10xx xxxx*/
			count += 2
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			/* 1110 xxxx  10xx xxxx  10xx xxxx*/
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytearr[count-2])
			char3 = uint16(bytearr[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", (count - 1)))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			/* 10xx xxxx,  1111 xxxx */
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	// The number of chars produced may be less than utflen
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}

/*
CONSTANT_String_info {
    u1 tag;
    u2 string_index;
}
*/
type ConstantStringInfo struct {
	StringIndex uint16
}

func readConstantStringInfo(reader *ClassReader) ConstantStringInfo {
	return ConstantStringInfo{
		StringIndex: reader.readUint16(),
	}
}

/*
CONSTANT_Class_info {
    u1 tag;
    u2 name_index;
}
*/
type ConstantClassInfo struct {
	NameIndex uint16
}

func readConstantClassInfo(reader *ClassReader) ConstantClassInfo {
	return ConstantClassInfo{
		NameIndex: reader.readUint16(),
	}
}

/*
CONSTANT_FieldRef_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_MethodRef_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
CONSTANT_InterfaceMethodRef_info {
    u1 tag;
    u2 class_index;
    u2 name_and_type_index;
}
*/
type ConstantMemberRefInfo struct {
	Tag              uint8
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func readConstantMemberRefInfo(reader *ClassReader, tag uint8) ConstantMemberRefInfo {
	return ConstantMemberRefInfo{
		Tag:              tag,
		ClassIndex:       reader.readUint16(),
		NameAndTypeIndex: reader.readUint16(),
	}
}

/*
CONSTANT_NameAndType_info {
    u1 tag;
    u2 name_index;
    u2 descriptor_index;
}
*/
type ConstantNameAndTypeInfo struct {
	NameIndex       uint16
	DescriptorIndex uint16
}

func readConstantNameAndTypeInfo(reader *ClassReader) ConstantNameAndTypeInfo {
	return ConstantNameAndTypeInfo{
		NameIndex:       reader.readUint16(),
		DescriptorIndex: reader.readUint16(),
	}
}

/******* invoke_dynamic *******/

/*
CONSTANT_InvokeDynamic_info {
    u1 tag;
    u2 bootstrap_method_attr_index;
    u2 name_and_type_index;
}
*/
type ConstantInvokeDynamicInfo struct {
	BootstrapMethodAttrIndex uint16
	NameAndTypeIndex         uint16
}

func readConstantInvokeDynamicInfo(reader *ClassReader) ConstantInvokeDynamicInfo {
	return ConstantInvokeDynamicInfo{
		BootstrapMethodAttrIndex: reader.readUint16(),
		NameAndTypeIndex:         reader.readUint16(),
	}
}

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
*/
type ConstantMethodHandleInfo struct {
	ReferenceKind  uint8
	ReferenceIndex uint16
}

func readConstantMethodHandleInfo(reader *ClassReader) ConstantMethodHandleInfo {
	return ConstantMethodHandleInfo{
		ReferenceKind:  reader.readUint8(),
		ReferenceIndex: reader.readUint16(),
	}
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
*/
type ConstantMethodTypeInfo struct {
	DescriptorIndex uint16
}

func readConstantMethodTypeInfo(reader *ClassReader) ConstantMethodTypeInfo {
	return ConstantMethodTypeInfo{
		DescriptorIndex: reader.readUint16(),
	}
}
