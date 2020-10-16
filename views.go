package tviewplus

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// NewReadOnlyTextView returns a readonly TextView that listens to changes of a StringHolder
func NewReadOnlyTextView(app *tview.Application, h *StringHolder) *tview.TextView {
	w := tview.NewTextView()
	w.SetText(h.value)
	h.AddDependent(func(old, new string) {
		app.QueueUpdateDraw(func() {
			w.SetText(new)
		})
	})
	return w
}

// NewInputView returns an InputField which listens to changes of a StringHolder
// and passes the focus to a FocusGroup when exiting the InputField
func NewInputView(f *FocusGroup, h *StringHolder) *tview.InputField {
	w := tview.NewInputField()
	w.SetText(h.value)
	f.Add(w)
	w.SetDoneFunc(func(k tcell.Key) {
		h.Set(w.GetText())
		f.HandleDone(w, k)
	})
	h.AddDependent(func(old, new string) {
		if w.GetText() != new {
			f.GetApplication().QueueUpdateDraw(func() {
				w.SetText(new)
			})
		}
	})
	return w
}

// NewDropDownView returns a DropDown which listens to changes of a StringListSelectionHolder
// and passes the focus to a FocusGroup when exiting the DropDown
func NewDropDownView(f *FocusGroup, h *StringListSelectionHolder) *tview.DropDown {
	w := tview.NewDropDown()
	w.SetOptions(h.list, func(text string, index int) {
		h.setSelection(SelectionWithIndex{Value: text, Index: index})
	})
	w.SetCurrentOption(h.Selection.Index) // -1 means no selection
	f.Add(w)
	w.SetDoneFunc(func(k tcell.Key) {
		f.HandleDone(w, k)
	})
	return w
}
