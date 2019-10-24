package rtda

import "github.com/taoyq1988/jvmgo/rtda/heap"

type LocalVars struct {
	slots []heap.Slot
}

func newLocalVars(size uint) LocalVars {
	slots := make([]heap.Slot, size)
	return LocalVars{slots: slots}
}

func (localVars *LocalVars) GetIntVar(idx uint) int32 {
	return localVars.GetLocalVar(idx).IntValue()
}

func (localVars *LocalVars) SetIntVar(idx uint, val int32) {
	localVars.SetLocalVar(idx, heap.NewIntSlot(val))
}

func (localVars *LocalVars) GetFloatVar(idx uint) float32 {
	return localVars.GetLocalVar(idx).FloatValue()
}

func (localVars *LocalVars) SetFloatVar(idx uint, val float32) {
	localVars.SetLocalVar(idx, heap.NewFloatSlot(val))
}

func (localVars *LocalVars) GetLongVar(idx uint) int64 {
	return localVars.GetLocalVar(idx).LongValue()
}

func (localVars *LocalVars) SetLongVar(idx uint, val int64) {
	localVars.SetLocalVar(idx, heap.NewLongSlot(val))
}

func (localVars *LocalVars) GetDoubleVar(idx uint) float64 {
	return localVars.GetLocalVar(idx).DoubleValue()
}

func (localVars *LocalVars) SetDoubleVar(idx uint, val float64) {
	localVars.SetLocalVar(idx, heap.NewDoubleSlot(val))
}

func (localVars *LocalVars) GetRefVar(idx uint) *heap.Object {
	return localVars.GetLocalVar(idx).Ref
}

func (localVars *LocalVars) SetRefVar(idx uint, val *heap.Object) {
	localVars.SetLocalVar(idx, heap.NewRefSlot(val))
}

func (localVars *LocalVars) GetLocalVar(idx uint) heap.Slot {
	return localVars.slots[idx]
}

func (localVars *LocalVars) SetLocalVar(idx uint, slot heap.Slot) {
	localVars.slots[idx] = slot
}

func (localVars *LocalVars) GetThis() *heap.Object {
	return localVars.GetRefVar(0)
}
