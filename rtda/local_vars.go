package rtda

type LocalVars struct {
	slots []Slot
}

func newLocalVars(size uint) LocalVars {
	slots := make([]Slot, 0)
	return LocalVars{slots: slots}
}
