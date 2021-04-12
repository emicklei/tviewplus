package tviewplus

type SelectionChangeHandler func(old, new SelectionWithIndex)

type ListChangeHandler func(old, new []string)

type StringListSelectionHolder struct {
	Selection      SelectionWithIndex
	list           []string
	dependents     []SelectionChangeHandler
	listDependents []ListChangeHandler
}

type SelectionWithIndex struct {
	Value string
	Index int
}

func (s *StringListSelectionHolder) Select(index int) {
	if index == -1 {
		s.setSelection(SelectionWithIndex{Index: -1})
		return
	}
	if index >= len(s.list) {
		return
	}
	s.setSelection(SelectionWithIndex{Value: s.list[index], Index: index})
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
	old := s.list
	// handlers may refer to the new list already
	s.list = newValue

	// reset selection because entry may be gone
	s.setSelection(SelectionWithIndex{Index: -1})

	// TODO this may be executed before selection updates...
	for _, each := range s.listDependents[:] {
		// call the handlers in a new go-routine
		// such that code exectued by the handler can update other widgets
		go each(old, newValue)
	}
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

func (s *StringListSelectionHolder) AddSelectionChangeDependent(h SelectionChangeHandler) {
	s.dependents = append(s.dependents, h)
}

func (s *StringListSelectionHolder) AddListChangeDependent(h ListChangeHandler) {
	s.listDependents = append(s.listDependents, h)
}
