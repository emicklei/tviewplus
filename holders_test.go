package tviewplus

import (
	"testing"
	"time"
)

func TestStringListSelectionHolder_Set(t *testing.T) {
	h := new(StringListSelectionHolder)
	var caught SelectionWithIndex
	h.AddDependent(func(old, new SelectionWithIndex) {
		caught = new
	})
	h.setSelection(SelectionWithIndex{Index: 1})
	// allow the go-routine to complete
	time.Sleep(10 * time.Millisecond)
	if caught.Index != 1 {
		t.Fail()
	}
}
