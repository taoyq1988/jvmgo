package rtda

import "github.com/taoyq1988/jvmgo/rtda/heap"

type LocalVars struct {
	slots []Slot
}

func newLocalVars(size uint) LocalVars {
	slots := make([]Slot, size)
	return LocalVars{slots: slots}
}

func (localVars *LocalVars) GetLocalVar(idx uint) Slot {
	return localVars.slots[idx]
}

func (localVars *LocalVars) SetLocalVar(idx uint, slot Slot) {
	localVars.slots[idx] = slot
}

func (localVars *LocalVars) GetIntVar(idx uint) int32 {
	return localVars.GetLocalVar(idx).IntValue()
}

func (localVars *LocalVars) SetIntVar(idx uint, val int32) {
	localVars.SetLocalVar(idx, heap.NewIntSlot(val))
}
