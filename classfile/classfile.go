package classfile

import (
	"fmt"
	"strconv"
)

const (
	magicNumber = 0xCAFEBABE
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
	ConstPool ConstantPool
	AccessFlags uint16
	ThisClass uint16
	SuperClass uint16
	Interfaces []uint16
	Fields []MemberInfo
	Methods []MemberInfo
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

func(cf *Classfile) parse(reader *ClassReader) (err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()

	cf.parseAndCheckMagic(reader)
	cf.parseAndCheckVersion(reader)
	cf.ConstPool = parseConstantPool(reader)
	cf.AccessFlags = reader.readUint16()
	cf.ThisClass = reader.readUint16()
	cf.SuperClass = reader.readUint16()
	cf.Interfaces = reader.readUint16s()
	cf.Fields = readMembers(reader)
	cf.Methods = readMembers(reader)
	cf.attributes = readAttributes(reader)
	return
}

func(cf *Classfile) parseAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != magicNumber {
		panic("magic number is not " + strconv.FormatInt(magicNumber, 16))
	}
}

func(cf *Classfile) parseAndCheckVersion(reader *ClassReader) {
	cf.MinorVersion = reader.readUint16()
	cf.MajorVersion = reader.readUint16()
	if cf.MajorVersion > supportVersion {
		panic("not support java version")
	}
}

