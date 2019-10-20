package heap

import (
	"github.com/taoyq1988/jvmgo/classfile"
)

type Constant interface {}
type ConstantPool []Constant

func newConstantPool(cf *classfile.Classfile) ConstantPool {
	cfCp := cf.ConstPool
	rtCp := make([]Constant, len(cfCp))
	for i:=1 ; i<len(cfCp); i++ {
		cpInfo := cfCp[i]
		switch x := cpInfo.(type) {
		case string:
			rtCp[i] = cpInfo
		case int32, float32:
			rtCp[i] = cpInfo
		case int64, float64:
			rtCp[i] = cpInfo
			i++
		case classfile.ConstantStringInfo:
			rtCp[i] = newConstantString(cf.GetUTF8(x.StringIndex))
		}
	}

	return rtCp
}

type ConstantString struct {
	goStr string
	jStr  *Object
}

func newConstantString(str string) *ConstantString {
	return &ConstantString{goStr: str}
}

func (s *ConstantString) GetJString() *Object {
	if s.jStr == nil {
		s.jStr = JSFromGoStr(s.goStr)
	}
	return s.jStr
}
