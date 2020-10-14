package tviewplus

type StringChangeHandler func(old, new string)
type SelectionChangeHandler func(old, new SelectionWithIndex)

type StringHolder struct {
	Value      string
	dependents []StringChangeHandler
}

func (s *StringHolder) AddDependent(h StringChangeHandler) {
	s.dependents = append(s.dependents, h)
}

func (s *StringHolder) Set(newValue string) {
	old := s.Value
	if newValue == old {
		return
	}
	s.Value = newValue
	for _, each := range s.dependents[:] {
		// call the handlers in a new go-routine
		// such that code exectued by the handler can update other widgets
		go each(old, newValue)
	}
}

type StringListSelectionHolder struct {
	Selection  SelectionWithIndex
	List       []string
	dependents []SelectionChangeHandler
}

type SelectionWithIndex struct {
	Value string
	Index int
}

func (s *StringListSelectionHolder) Set(newValue []string) {
	if len(s.List) == len(newValue) {
		same := true
		for i, each := range s.List {
			if same && newValue[i] != each {
				same = false
			}
		}
		if same {
			return
		}
	}
	s.List = newValue
	s.setSelection(SelectionWithIndex{Index: -1})
}

func (s *StringListSelectionHolder) setSelection(sel SelectionWithIndex) {
	old := s.Selection
	for _, each := range s.dependents[:] {
		// call the handlers in a new go-routine
		// such that code exectued by the handler can update other widgets
		go each(old, sel)
	}
	s.Selection = sel
}

func (s *StringListSelectionHolder) AddDependent(h SelectionChangeHandler) {
	s.dependents = append(s.dependents, h)
}

type WriterStringHolderAdaptor struct {
	target *StringHolder
}

// Write is part of io.Writer
func (w WriterStringHolderAdaptor) Write(data []byte) (int, error) {
	w.target.Set(w.target.Value + string(data))
	return len(data), nil
}
