package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func describe(describe string, objFocusReturn tview.Primitive, pageInvoke string, shortcuts string, titlePageInvoke string) {
	describePod.SetText(describe)

	stringShortcuts = SHORTCUTS_DESCRIBE
	verifyContext()

	pages.SwitchToPage("DescribePod")
	pages.SetTitle("Describe")

	app.SetFocus(describePod)

	describePodShortcuts(objFocusReturn, pageInvoke, shortcuts, titlePageInvoke)

}

func describePodShortcuts(objFocusReturn tview.Primitive, switchToPage string, shortcuts string, titlePageInvoke string) {
	describePod.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlM) {
			pages.SwitchToPage(switchToPage)
			pages.SetTitle(titlePageInvoke)

			stringShortcuts = shortcuts
			verifyContext()

			app.SetFocus(objFocusReturn)
		}
		return event
	})
}
