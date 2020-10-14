package main

import (
	"fmt"
	"log"

	"github.com/emicklei/tviewplus"
	"github.com/rivo/tview"
)

// go run *.go 2> stderr.log
// tail -f stderr.log
func main() {
	bin := &Bindings{
		Name:    new(tviewplus.StringHolder),
		List:    new(tviewplus.StringListSelectionHolder),
		Console: new(tviewplus.StringHolder),
	}
	bin.Name.Set("initial")
	bin.List.Set([]string{"initial", "selection"})

	bin.Name.AddDependent(func(old, new string) {
		bin.Console.Set(fmt.Sprintf("Name changed from [%s] to [%s]", old, new))
	})
	bin.List.AddDependent(func(old, new tviewplus.SelectionWithIndex) {
		bin.Console.Set(fmt.Sprintf("Dropdown selection changed from [%v] to [%v]", old, new))
	})

	app := tview.NewApplication()
	foc := tviewplus.NewFocusGroup(app)

	flex := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tviewplus.NewStaticView(" [yellow]InputView"), 1, 1, false).
		AddItem(tviewplus.NewInputView(foc, bin.Name), 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(tviewplus.NewStaticView(" [red]DropDownView"), 1, 1, false).
		AddItem(tviewplus.NewDropDownView(foc, bin.List), 1, 1, false).
		AddItem(tview.NewBox().SetBorderPadding(1, 0, 0, 0), 1, 1, false).
		AddItem(tviewplus.NewStaticView(" [gray]TextView"), 1, 1, false).
		AddItem(tviewplus.NewTextView(app, bin.Console), 1, 1, false)

	if err := app.SetRoot(flex, true).SetFocus(foc.GetFocus()).EnableMouse(true).Run(); err != nil {
		log.Println(err)
	}
}

type Bindings struct {
	Name    *tviewplus.StringHolder
	List    *tviewplus.StringListSelectionHolder
	Console *tviewplus.StringHolder
}
