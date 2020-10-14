package tviewplus

import "testing"

func TestStringListSelectionHolder_Set(t *testing.T) {
	h := new(StringListSelectionHolder)
	var caught SelectionWithIndex
	h.AddDependent(func(old, new SelectionWithIndex) {
		caught = new
	})
	h.
}
