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
	numBootstrapMethods := reader.readUint16()
	bootstrapMethods := make([]BootstrapMethod, numBootstrapMethods)
	for i := range bootstrapMethods {
		bootstrapMethods[i] = BootstrapMethod{
			BootstrapMethodRef: reader.readUint16(),
			BootstrapArguments: reader.readUint16s(),
		}
	}
	return BootstrapMethodsAttribute{
		BootstrapMethods: bootstrapMethods,
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
		MaxStack:       reader.readUint16(),
		MaxLocals:      reader.readUint16(),
		Code:           reader.readBytes(reader.readUint32()),
		ExceptionTable: readExceptionTable(reader),
		AttributeTable: AttributeTable{
			readAttributes(reader),
		},
	}
}

type ExceptionTableEntry struct {
	StartPc   uint16
	EndPc     uint16
	HandlerPc uint16
	CatchType uint16
}

func readExceptionTable(reader *ClassReader) []ExceptionTableEntry {
	tableLength := reader.readUint16()
	exceptionTable := make([]ExceptionTableEntry, tableLength)
	for i := range exceptionTable {
		exceptionTable[i] = ExceptionTableEntry{
			StartPc:   reader.readUint16(),
			EndPc:     reader.readUint16(),
			HandlerPc: reader.readUint16(),
			CatchType: reader.readUint16(),
		}
	}
	return exceptionTable
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
		ConstantValueIndex: reader.readUint16(),
	}
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
		ClassIndex:  reader.readUint16(),
		MethodIndex: reader.readUint16(),
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
	numberOfClasses := reader.readUint16()
	classes := make([]InnerClassInfo, numberOfClasses)
	for i := range classes {
		classes[i] = InnerClassInfo{
			InnerClassInfoIndex:   reader.readUint16(),
			OuterClassInfoIndex:   reader.readUint16(),
			InnerNameIndex:        reader.readUint16(),
			InnerClassAccessFlags: reader.readUint16(),
		}
	}
	return InnerClassesAttribute{
		Classes: classes,
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
	tableLength := reader.readUint16()
	lineNumberTable := make([]LineNumberTableEntry, tableLength)
	for i := range lineNumberTable {
		lineNumberTable[i] = LineNumberTableEntry{
			StartPC:    reader.readUint16(),
			LineNumber: reader.readUint16(),
		}
	}
	return LineNumberTableAttribute{
		LineNumberTable: lineNumberTable,
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
	tableLength := reader.readUint16()
	localVariableTable := make([]LocalVariableTableEntry, tableLength)
	for i := range localVariableTable {
		localVariableTable[i] = LocalVariableTableEntry{
			StartPc:         reader.readUint16(),
			Length:          reader.readUint16(),
			NameIndex:       reader.readUint16(),
			DescriptorIndex: reader.readUint16(),
			Index:           reader.readUint16(),
		}
	}
	return LocalVariableTableAttribute{
		LocalVariableTable: localVariableTable,
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
	tableLength := reader.readUint16()
	localVariableTypeTable := make([]LocalVariableTypeTableEntry, tableLength)
	for i := range localVariableTypeTable {
		localVariableTypeTable[i] = LocalVariableTypeTableEntry{
			StartPc:        reader.readUint16(),
			Length:         reader.readUint16(),
			NameIndex:      reader.readUint16(),
			SignatureIndex: reader.readUint16(),
			Index:          reader.readUint16(),
		}
	}
	return LocalVariableTypeTableAttribute{
		LocalVariableTypeTable: localVariableTypeTable,
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
		SignatureIndex: reader.readUint16(),
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
	return SourceFileAttribute{SourceFileIndex: reader.readUint16()}
}
