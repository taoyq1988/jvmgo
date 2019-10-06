package rtda

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
