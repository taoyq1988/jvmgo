package rtda

import "github.com/taoyq1988/jvmgo/rtda/heap"

var (
	_shimClass = &heap.Class{Name: "~shim"}

	shimReturnMethod = &heap.Method{
		ClassMember: newShimMember("<return>"),
		MethodData: heap.MethodData{
			Code: []byte{0xb1}, // return
		},
	}

	shimAThrowMethod = &heap.Method{
		ClassMember: newShimMember("<athrow>"),
		MethodData: heap.MethodData{
			Code: []byte{0xbf}, // athrow
		},
	}

	ShimBootstrapMethod = &heap.Method{
		ClassMember: newShimMember("<bootstrap>"),
		MethodData: heap.MethodData{
			Code:      []byte{0xff, 0xb1}, // bootstrap, return
			MaxStack:  8,
			MaxLocals: 8,
		},
		ParamSlotCount: 2,
	}
)

func newShimMember(name string) heap.ClassMember {
	return heap.ClassMember{
		AccessFlag: heap.AccStatic,
		Name:       name,
		Class:      _shimClass,
	}
}
