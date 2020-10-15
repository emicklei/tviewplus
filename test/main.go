package main

import (
	"fmt"
	"log"

	"github.com/emicklei/tviewplus"
	"github.com/rivo/tview"
)

// go run *.go 2> stderr.log
// tail -f stderr.log

type Bindings struct {
	Name    *tviewplus.StringHolder
	List    *tviewplus.StringListSelectionHolder
	Console *tviewplus.StringHolder
}

func main() {
	bin := &Bindings{
		Name:    new(tviewplus.StringHolder),
		List:    new(tviewplus.StringListSelectionHolder),
		Console: new(tviewplus.StringHolder),
	}

	// initial values
	bin.Name.Set("edit me")
	bin.List.Set([]string{" choice A ", " choice B ", " choice C "})
	bin.List.Select(0)
	bin.Console.Append("Cycle through editable views using Tab,Enter,Escape,Back Tab\n")

	// inter view dependencies
	bin.Name.AddDependent(func(old, new string) {
		bin.Console.Append(fmt.Sprintf("Name changed from [%s] to [%s]\n", old, new))
	})

	bin.List.AddDependent(func(old, new tviewplus.SelectionWithIndex) {
		bin.Console.Append(fmt.Sprintf("Dropdown selection changed from [%v] to [%v]\n", old, new))
		bin.Name.Set(new.Value)
	})

	// compose the app
	app := tview.NewApplication()

	// for cycling through editable views
	foc := tviewplus.NewFocusGroup(app)

	// editor for Name
	nameField := tviewplus.NewInputView(foc, bin.Name)

	// editor for List
	choiceDropdown := tviewplus.NewDropDownView(foc, bin.List)

	// viewer for Console
	console := tviewplus.NewTextView(app, bin.Console)
	console.SetBorder(true).SetTitle("console")

	// layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tviewplus.NewStaticView(" [gray]InputView"), 1, 1, false).
		AddItem(nameField, 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(tviewplus.NewStaticView(" [gray]DropDownView"), 1, 1, false).
		AddItem(choiceDropdown, 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(tviewplus.NewStaticView(" [gray]TextView"), 1, 1, false).
		AddItem(console, 10, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(foc.GetFocus()).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}
