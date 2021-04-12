package tviewplus

import (
	"testing"
	"time"
)

func TestStringListSelectionHolder_SetSelection(t *testing.T) {
	h := new(StringListSelectionHolder)
	var caught SelectionWithIndex
	h.AddSelectionChangeDependent(func(old, new SelectionWithIndex) {
		caught = new
	})
	h.setSelection(SelectionWithIndex{Index: 1})
	// allow the go-routine to complete
	time.Sleep(10 * time.Millisecond)
	if caught.Index != 1 {
		t.Fail()
	}
}

func TestStringListSelectionHolder_Set(t *testing.T) {
	h := new(StringListSelectionHolder)
	var caught []string
	h.AddListChangeDependent(func(old, new []string) {
		caught = new
	})
	h.Set([]string{"hello"})
	// allow the go-routine to complete
	time.Sleep(10 * time.Millisecond)
	if caught[0] != "hello" {
		t.Fail()
	}
}
