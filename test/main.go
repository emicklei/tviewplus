package main

import (
	"fmt"
	"log"

	"github.com/emicklei/tviewplus"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// go run *.go 2> stderr.log
// tail -f stderr.log

type Bindings struct {
	Name    *tviewplus.StringHolder
	List    *tviewplus.StringListSelectionHolder
	Choices *tviewplus.StringListSelectionHolder
	Checked *tviewplus.BoolHolder
	Console *tviewplus.StringHolder
}

func main() {
	bin := &Bindings{
		Name:    new(tviewplus.StringHolder),
		List:    new(tviewplus.StringListSelectionHolder),
		Choices: new(tviewplus.StringListSelectionHolder),
		Checked: new(tviewplus.BoolHolder),
		Console: new(tviewplus.StringHolder),
	}

	// initial values
	bin.Name.Set("edit me")
	bin.Choices.Set([]string{" choice A ", " choice B ", " choice C "})
	bin.Choices.Select(0)
	bin.Console.Append("Cycle through editable views using Tab,Enter,Escape,Back Tab\n")
	bin.List.Set([]string{"item 1", "item 2", "item 3"})

	// inter view dependencies
	bin.Name.AddDependent(func(old, new string) {
		bin.Console.Append(fmt.Sprintf("Name changed from [%s] to [%s]\n", old, new))
	})

	bin.Choices.AddSelectionChangeDependent(func(old, new tviewplus.SelectionWithIndex) {
		bin.Console.Append(fmt.Sprintf("Dropdown selection changed from [%v] to [%v]\n", old, new))
		bin.Name.Set(new.Value)
	})

	bin.List.AddSelectionChangeDependent(func(old, new tviewplus.SelectionWithIndex) {
		bin.Console.Append(fmt.Sprintf("List selection changed from [%v] to [%v]\n", old, new))
	})

	bin.Checked.AddDependent(func(old, new bool) {
		bin.Console.Append(fmt.Sprintf("Checkbox clicked:%v\n", new))
	})

	// compose the app
	app := tview.NewApplication()

	// for cycling through editable views
	foc := tviewplus.NewFocusGroup(app)

	// editor for Name
	nameField := tviewplus.NewInputView(foc, bin.Name).SetFieldWidth(16)
	nameFieldLabel := tview.NewTextView().SetDynamicColors(true).SetText(" [gray]InputField with StringHolder")

	// editor for DropDown
	choiceDropdown := tviewplus.NewDropDownView(foc, bin.Choices)
	choiceDropdown.SetTextOptions("", "", "", "▼", "---")
	choiceDropdownLabel := tview.NewTextView().SetDynamicColors(true).SetText(" [gray]DropDown with StringListSelectionHolder")

	// readonly List
	showList := tviewplus.NewListView(foc, bin.List)
	showListLabel := tview.NewTextView().SetDynamicColors(true).SetText(" [gray]List with StringListSelectionHolder")

	// button
	button := tviewplus.NewButtonView(foc).SetLabel("OK")
	button.SetSelectedFunc(func() {
		bin.Console.Append(fmt.Sprintf("Button OK clicked\n"))
	})
	buttonLabel := tview.NewTextView().SetDynamicColors(true).SetText(" [gray]Button")

	// checkbox ✓
	checkbox := tviewplus.NewCheckboxView(foc, bin.Checked).SetLabel("Tick me ")
	checkbox.SetCheckedString("✓")
	checkboxLabel := tview.NewTextView().SetDynamicColors(true).SetText(" [gray]Checkbox with BoolHolder")

	// viewer for Console
	console := tviewplus.NewReadOnlyTextView(app, bin.Console)
	console.SetBorder(true).SetTitle("log")
	consoleLabel := tview.NewTextView().SetDynamicColors(true).SetText(" [gray]TextView with StringHolder")

	// layout
	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(nameFieldLabel, 1, 1, false).
		AddItem(nameField, 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(choiceDropdownLabel, 1, 1, false).
		AddItem(choiceDropdown, 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(showListLabel, 1, 1, false).
		AddItem(showList, 4, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(buttonLabel, 1, 1, false).
		AddItem(button, 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(checkboxLabel, 1, 1, false).
		AddItem(checkbox, 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(consoleLabel, 1, 1, false).
		AddItem(console, 10, 1, false)

	tview.Styles.PrimaryTextColor = tcell.ColorWhite
	tview.Styles.ContrastBackgroundColor = tcell.ColorBlue
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack

	if err := app.SetRoot(flex, true).SetFocus(foc.GetFocus()).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}
