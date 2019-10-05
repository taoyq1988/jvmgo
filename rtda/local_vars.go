package rtda

type LocalVars struct {
	slots []Slot
}

func newLocalVars(size uint) LocalVars {
	slots := make([]Slot, size)
	return LocalVars{slots: slots}
}
