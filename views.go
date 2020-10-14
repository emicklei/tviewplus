package tviewplus

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

func NewStaticView(label string) *tview.TextView {
	w := tview.NewTextView()
	w.SetDynamicColors(true)
	w.SetText(label)
	return w
}

func NewTextView(app *tview.Application, h *StringHolder) *tview.TextView {
	w := tview.NewTextView()
	h.AddDependent(func(old, new string) {
		app.QueueUpdateDraw(func() {
			w.SetText(new)
		})
	})
	return w
}

func NewInputView(f *FocusGroup, h *StringHolder) *tview.InputField {
	w := tview.NewInputField()
	w.SetText(h.Value)
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

func NewDropDownView(f *FocusGroup, h *StringListSelectionHolder) *tview.DropDown {
	w := tview.NewDropDown()
	w.SetOptions(h.List, func(text string, index int) {
		h.setSelection(SelectionWithIndex{Value: text, Index: index})
	})
	f.Add(w)
	w.SetDoneFunc(func(k tcell.Key) {
		f.HandleDone(w, k)
	})
	return w
}
