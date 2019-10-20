package classfile

import (
	"fmt"
	"strconv"
)

const (
	magicNumber    = 0xCAFEBABE
	supportVersion = 52
)

/*
ClassFile {
    u4             magic;
    u2             minor_version;
    u2             major_version;
    u2             constant_pool_count;
    cp_info        constant_pool[constant_pool_count-1];
    u2             access_flags;
    u2             this_class;
    u2             super_class;
    u2             interfaces_count;
    u2             interfaces[interfaces_count];
    u2             fields_count;
    field_info     fields[fields_count];
    u2             methods_count;
    method_info    methods[methods_count];
    u2             attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type Classfile struct {
	//Magic uint32
	MinorVersion uint16
	MajorVersion uint16
	ConstPool    ConstantPool
	AccessFlags  uint16
	ThisClass    uint16
	SuperClass   uint16
	Interfaces   []uint16
	Fields       []MemberInfo
	Methods      []MemberInfo
	AttributeTable
}

func Parse(data []byte) (*Classfile, error) {
	reader := newClassReader(data)
	cf := &Classfile{}
	err := cf.parse(reader)
	if err != nil {
		return nil, err
	}
	return cf, nil
}

func (cf *Classfile) parse(reader *ClassReader) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	reader.cf = cf
	cf.parseAndCheckMagic(reader)
	cf.parseAndCheckVersion(reader)
	cf.ConstPool = parseConstantPool(reader)
	cf.AccessFlags = reader.ReadUint16()
	cf.ThisClass = reader.ReadUint16()
	cf.SuperClass = reader.ReadUint16()
	cf.Interfaces = reader.readUint16s()
	cf.Fields = readMembers(reader)
	cf.Methods = readMembers(reader)
	cf.AttributeTable = readAttributes(reader)
	return
}

func (cf *Classfile) parseAndCheckMagic(reader *ClassReader) {
	magic := reader.ReadUint32()
	if magic != magicNumber {
		panic("magic number is not " + strconv.FormatInt(magicNumber, 16))
	}
}

func (cf *Classfile) parseAndCheckVersion(reader *ClassReader) {
	cf.MinorVersion = reader.ReadUint16()
	cf.MajorVersion = reader.ReadUint16()
	if cf.MajorVersion > supportVersion {
		panic("not support java version")
	}
}

func (cf *Classfile) GetThisClassName() string {
	return cf.GetClassNameOf(cf.ThisClass)
}

func (cf *Classfile) GetSuperClassName() string {
	return cf.GetClassNameOf(cf.SuperClass)
}

func (cf *Classfile) GetInterfaceNames() []string {
	interfaces := make([]string, len(cf.Interfaces))
	for i, index := range cf.Interfaces {
		interfaces[i] = cf.GetClassNameOf(index)
	}
	return interfaces
}

func (cf *Classfile) GetConstantInfo(index uint16) ConstantInfo {
	return cf.getConstantInfo(index)
}

func (cf *Classfile) GetUTF8(index uint16) string {
	return cf.getUtf8(index)
}

func (cf *Classfile) GetClassNameOf(index uint16) string {
	if index == 0 {
		return ""
	}
	classInfo := cf.getConstantInfo(index).(ConstantClassInfo)
	return cf.getUtf8(classInfo.NameIndex)
}

func (cf *Classfile) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := cf.ConstPool[index]; cpInfo == nil {
		panic(fmt.Errorf("invalid constant pool index: %d", index))
	} else {
		return cpInfo
	}
}

func (cf *Classfile) getUtf8(index uint16) string {
	if index == 0 {
		return ""
	}
	return cf.getConstantInfo(index).(string)
}
