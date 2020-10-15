package tviewplus

type SelectionChangeHandler func(old, new SelectionWithIndex)

type StringListSelectionHolder struct {
	Selection  SelectionWithIndex
	list       []string
	dependents []SelectionChangeHandler
}

type SelectionWithIndex struct {
	Value string
	Index int
}

func (s *StringListSelectionHolder) Set(newValue []string) {
	if len(s.list) == len(newValue) {
		same := true
		for i, each := range s.list {
			if same && newValue[i] != each {
				same = false
			}
		}
		if same {
			return
		}
	}
	s.list = newValue
	s.setSelection(SelectionWithIndex{Index: -1})
}

func (s *StringListSelectionHolder) setSelection(sel SelectionWithIndex) {
	if s.Selection.Index == sel.Index {
		return
	}
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
