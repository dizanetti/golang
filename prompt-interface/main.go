package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

func main() {
	initProg()

	flex := tview.NewFlex().
		AddItem(list, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(infoPages, 5, 2, false).
			AddItem(pages, 0, 1, false), 0, 3, false)

	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

func initProg() {
	configureShortcuts()

	createTableContext()
	createTablePods(GET_PODS)

	setPages()

	verifyContext()

	schedulerSeconds(2, verifyContext)
}

func configureShortcuts() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 121 {
			app.SetFocus(list)
		} else if event.Rune() == 113 {
			app.Stop()
		}

		return event
	})
}
