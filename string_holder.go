package tviewplus

type StringChangeHandler func(old, new string)

// StringHolder has a string Value that notifies all its dependents when the value changes.
type StringHolder struct {
	value      string
	dependents []StringChangeHandler
}

// Write is part of io.Writer
func (s *StringHolder) Write(data []byte) (int, error) {
	s.Append(string(data))
	return len(data), nil
}

func (s *StringHolder) AddDependent(h StringChangeHandler) {
	s.dependents = append(s.dependents, h)
}

func (s *StringHolder) Append(moreValue string) {
	s.Set(s.value + moreValue)
}

func (s *StringHolder) Get() string {
	return s.value
}

func (s *StringHolder) Set(newValue string) {
	old := s.value
	if newValue == old {
		return
	}
	s.value = newValue
	for _, each := range s.dependents[:] {
		// call the handlers in a new go-routine
		// such that code exectued by the handler can update other widgets
		go each(old, newValue)
	}
}
