package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func load(podName string, typeOfService string, objFocusReturn tview.Primitive, pageInvoke string) {
	var outputTypeLoad string
	if settings.DefaultOutputFormatted == 0 {
		outputTypeLoad = "yaml"
	} else {
		outputTypeLoad = "json"
	}

	load, errLoad := execute("kubectl get " + typeOfService + " -o=" + outputTypeLoad + " " + podName)
	if errLoad != nil {
		informationText.SetText(errLoad.Error()).SetTextColor(tcell.ColorRed)
	} else {
		loadConfiguration.SetText(strings.Join(load, "\n"))

		stringShortcuts = SHORTCUTS_LOAD_CONFIG
		verifyContext()

		pages.SwitchToPage("LoadConfiguration")
		pages.SetTitle("Load Configuration")

		app.SetFocus(loadConfiguration)

		loadConfigurationShortcuts(objFocusReturn, pageInvoke, SHORTCUTS_LOAD_CONFIG)
	}
}

func loadConfigurationShortcuts(objFocusReturn tview.Primitive, switchToPage string, shortcuts string) {
	loadConfiguration.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlX) {
			pages.SwitchToPage(switchToPage)

			stringShortcuts = shortcuts
			verifyContext()

			app.SetFocus(objFocusReturn)
		}
		return event
	})
}
