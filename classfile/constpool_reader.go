package classfile

import (
	"github.com/taoyq1988/jvmgo/vmutils"
	"math"
)

/*
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantIntegerInfo(reader *ClassReader) int32 {
	return int32(reader.ReadUint32())
}

/*
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
*/
func readConstantFloatInfo(reader *ClassReader) float32 {
	return math.Float32frombits(reader.ReadUint32())
}

/*
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantLongInfo(reader *ClassReader) int64 {
	return int64(reader.ReadUint64())
}

/*
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
func readConstantDoubleInfo(reader *ClassReader) float64 {
	return math.Float64frombits(reader.ReadUint64())
}

/*
CONSTANT_Utf8_info {
    u1 tag;
    u2 length;
    u1 bytes[length];
}
*/
func readConstantUtf8Info(reader *ClassReader) string {
	length := int(reader.ReadUint16())
	bytes := reader.ReadBytes(length)
	return vmutils.DecodeMUTF8(bytes)
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
		StringIndex: reader.ReadUint16(),
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
		NameIndex: reader.ReadUint16(),
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
type ConstantFieldRefInfo constantMemberRefInfo
type ConstantMethodRefInfo constantMemberRefInfo
type ConstantInterfaceMethodRefInfo constantMemberRefInfo

type constantMemberRefInfo struct {
	ClassIndex       uint16
	NameAndTypeIndex uint16
}

func readConstantFieldRefInfo(reader *ClassReader) ConstantFieldRefInfo {
	return ConstantFieldRefInfo(readConstantMemberRefInfo(reader))
}
func readConstantMethodRefInfo(reader *ClassReader) ConstantMethodRefInfo {
	return ConstantMethodRefInfo(readConstantMemberRefInfo(reader))
}
func readConstantInterfaceMethodRefInfo(reader *ClassReader) ConstantInterfaceMethodRefInfo {
	return ConstantInterfaceMethodRefInfo(readConstantMemberRefInfo(reader))
}

func readConstantMemberRefInfo(reader *ClassReader) constantMemberRefInfo {
	return constantMemberRefInfo{
		ClassIndex:       reader.ReadUint16(),
		NameAndTypeIndex: reader.ReadUint16(),
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
		NameIndex:       reader.ReadUint16(),
		DescriptorIndex: reader.ReadUint16(),
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
		BootstrapMethodAttrIndex: reader.ReadUint16(),
		NameAndTypeIndex:         reader.ReadUint16(),
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
		ReferenceKind:  reader.ReadUint8(),
		ReferenceIndex: reader.ReadUint16(),
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
		DescriptorIndex: reader.ReadUint16(),
	}
}
