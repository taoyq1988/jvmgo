package classfile

/*
BootstrapMethods_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 num_bootstrap_methods;
    {   u2 bootstrap_method_ref;
        u2 num_bootstrap_arguments;
        u2 bootstrap_arguments[num_bootstrap_arguments];
    } bootstrap_methods[num_bootstrap_methods];
}
*/
type BootstrapMethodsAttribute struct {
	BootstrapMethods []BootstrapMethod
}

type BootstrapMethod struct {
	BootstrapMethodRef uint16
	BootstrapArguments []uint16
}

func readBootstrapMethodsAttribute(reader *ClassReader) BootstrapMethodsAttribute {
	return BootstrapMethodsAttribute{
		BootstrapMethods: reader.readTable(func(reader *ClassReader) BootstrapMethod {
			return BootstrapMethod{
				BootstrapMethodRef: reader.ReadUint16(),
				BootstrapArguments: reader.readUint16s(),
			}
		}).([]BootstrapMethod),
	}
}

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	MaxStack       uint16
	MaxLocals      uint16
	Code           []byte
	ExceptionTable []ExceptionTableEntry
	AttributeTable
}

func readCodeAttribute(reader *ClassReader) CodeAttribute {
	return CodeAttribute{
		MaxStack:       reader.ReadUint16(),
		MaxLocals:      reader.ReadUint16(),
		Code:           reader.ReadBytes(int(reader.ReadUint32())),
		ExceptionTable: readExceptionTable(reader),
		AttributeTable: readAttributes(reader),
	}
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func readExceptionTable(reader *ClassReader) []ExceptionTableEntry {
	return reader.readTable(func(reader *ClassReader) ExceptionTableEntry {
		return ExceptionTableEntry{
			StartPc:   reader.ReadUint16(),
			EndPc:     reader.ReadUint16(),
			HandlerPc: reader.ReadUint16(),
			CatchType: reader.ReadUint16(),
		}
	}).([]ExceptionTableEntry)
}

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
*/
type ConstantValueAttribute struct {
	ConstantValueIndex uint16
}

func readConstantValueAttribute(reader *ClassReader) ConstantValueAttribute {
	return ConstantValueAttribute{
		ConstantValueIndex: reader.ReadUint16(),
	}
}

/*
EnclosingMethod_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 class_index;
    u2 method_index;
}
*/
type EnclosingMethodAttribute struct {
	ClassIndex  uint16
	MethodIndex uint16
}

func readEnclosingMethodAttribute(reader *ClassReader) EnclosingMethodAttribute {
	return EnclosingMethodAttribute{
		ClassIndex:  reader.ReadUint16(),
		MethodIndex: reader.ReadUint16(),
	}
}

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
*/
type ExceptionsAttribute struct {
	ExceptionIndexTable []uint16
}

func readExceptionsAttribute(reader *ClassReader) ExceptionsAttribute {
	return ExceptionsAttribute{
		ExceptionIndexTable: reader.readUint16s(),
	}
}

/*
InnerClasses_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_classes;
    {   u2 inner_class_info_index;
        u2 outer_class_info_index;
        u2 inner_name_index;
        u2 inner_class_access_flags;
    } classes[number_of_classes];
}
*/
type InnerClassesAttribute struct {
	Classes []InnerClassInfo
}

type InnerClassInfo struct {
	InnerClassInfoIndex   uint16
	OuterClassInfoIndex   uint16
	InnerNameIndex        uint16
	InnerClassAccessFlags uint16
}

func readInnerClassesAttribute(reader *ClassReader) InnerClassesAttribute {
	return InnerClassesAttribute{
		Classes: reader.readTable(func(reader *ClassReader) InnerClassInfo {
			return InnerClassInfo{
				InnerClassInfoIndex:   reader.ReadUint16(),
				OuterClassInfoIndex:   reader.ReadUint16(),
				InnerNameIndex:        reader.ReadUint16(),
				InnerClassAccessFlags: reader.ReadUint16(),
			}
		}).([]InnerClassInfo),
	}
}

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	LineNumberTable []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	StartPC    uint16
	LineNumber uint16
}

func readLineNumberTableAttribute(reader *ClassReader) LineNumberTableAttribute {
	return LineNumberTableAttribute{
		LineNumberTable: reader.readTable(func(reader *ClassReader) LineNumberTableEntry {
			return LineNumberTableEntry{
				StartPC:    reader.ReadUint16(),
				LineNumber: reader.ReadUint16(),
			}
		}).([]LineNumberTableEntry),
	}
}

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
*/
type LocalVariableTableAttribute struct {
	LocalVariableTable []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	StartPc         uint16
	Length          uint16
	NameIndex       uint16
	DescriptorIndex uint16
	Index           uint16
}

func readLocalVariableTableAttribute(reader *ClassReader) LocalVariableTableAttribute {
	return LocalVariableTableAttribute{
		LocalVariableTable: reader.readTable(func(reader *ClassReader) LocalVariableTableEntry {
			return LocalVariableTableEntry{
				StartPc:         reader.ReadUint16(),
				Length:          reader.ReadUint16(),
				NameIndex:       reader.ReadUint16(),
				DescriptorIndex: reader.ReadUint16(),
				Index:           reader.ReadUint16(),
			}
		}).([]LocalVariableTableEntry),
	}
}

/*
LocalVariableTypeTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_type_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 signature_index;
        u2 index;
    } local_variable_type_table[local_variable_type_table_length];
}
*/
type LocalVariableTypeTableAttribute struct {
	LocalVariableTypeTable []LocalVariableTypeTableEntry
}

type LocalVariableTypeTableEntry struct {
	StartPc        uint16
	Length         uint16
	NameIndex      uint16
	SignatureIndex uint16
	Index          uint16
}

func readLocalVariableTypeTableAttribute(reader *ClassReader) LocalVariableTypeTableAttribute {
	return LocalVariableTypeTableAttribute{
		LocalVariableTypeTable: reader.readTable(func(reader *ClassReader) LocalVariableTypeTableEntry {
			return LocalVariableTypeTableEntry{
				StartPc:        reader.ReadUint16(),
				Length:         reader.ReadUint16(),
				NameIndex:      reader.ReadUint16(),
				SignatureIndex: reader.ReadUint16(),
				Index:          reader.ReadUint16(),
			}
		}).([]LocalVariableTypeTableEntry),
	}
}

/*
Signature_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 signature_index;
}
*/
type SignatureAttribute struct {
	SignatureIndex uint16
}

func readSignatureAttribute(reader *ClassReader) SignatureAttribute {
	return SignatureAttribute{
		SignatureIndex: reader.ReadUint16(),
	}
}

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
*/
type SourceFileAttribute struct {
	SourceFileIndex uint16
}

func readSourceFileAttribute(reader *ClassReader) SourceFileAttribute {
	return SourceFileAttribute{SourceFileIndex: reader.ReadUint16()}
}

/******** markers ********/

type MarkerAttribute struct{}

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
*/
type SyntheticAttribute struct {
	MarkerAttribute
}
