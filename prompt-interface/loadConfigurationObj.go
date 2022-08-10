package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func load(podName string, typeOfService string, objFocusReturn tview.Primitive, pageInvoke string, shortcuts string, titlePageInvoke string) {
	var outputTypeLoad string
	if settings.DefaultOutputFormatted == 0 {
		outputTypeLoad = "yaml"
	} else {
		outputTypeLoad = "json"
	}

	load, errLoad := execute("kubectl get " + typeOfService + " -o=" + outputTypeLoad + " " + podName)
	if errLoad != nil {
		FooterinformationText.SetText(errLoad.Error()).SetTextColor(tcell.ColorRed)
	} else {
		loadConfiguration.SetText(strings.Join(load, "\n"))

		stringShortcuts = SHORTCUTS_LOAD_CONFIG
		verifyContext()

		pages.SwitchToPage("LoadConfiguration")
		pages.SetTitle("Load Configuration")

		app.SetFocus(loadConfiguration)

		loadConfigurationShortcuts(objFocusReturn, pageInvoke, shortcuts, titlePageInvoke)
	}
}

func loadConfigurationShortcuts(objFocusReturn tview.Primitive, switchToPage string, shortcuts string, titlePageInvoke string) {
	loadConfiguration.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlX) {
			pages.SwitchToPage(switchToPage)
			pages.SetTitle(titlePageInvoke)

			stringShortcuts = shortcuts
			verifyContext()

			app.SetFocus(objFocusReturn)
		}
		return event
	})
}
