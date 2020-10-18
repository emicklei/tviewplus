package tviewplus

type BoolChangeHandler func(old, new bool)

type BoolHolder struct {
	value      bool
	dependents []BoolChangeHandler
}

func (b *BoolHolder) AddDependent(h BoolChangeHandler) {
	b.dependents = append(b.dependents, h)
}

func (b *BoolHolder) Get() bool {
	return b.value
}

func (b *BoolHolder) Set(newValue bool) {
	old := b.value
	if newValue == old {
		return
	}
	b.value = newValue
	for _, each := range b.dependents[:] {
		// call the handlers in a new go-routine
		// such that code exectued by the handler can update other widgets
		go each(old, newValue)
	}
}
