package main

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var app = tview.NewApplication()

var settings AppSettings

func main() {
	initProg()

	flex := tview.NewFlex().
		AddItem(pagesMenu, 0, 1, true).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(infoPages, 5, 2, false).
			AddItem(pages, 0, 1, false).
			AddItem(FooterPages, 4, 2, false), 0, 3, false)

	if err := app.SetRoot(flex, true).EnableMouse(false).Run(); err != nil {
		panic(err)
	}
}

func initProg() {
	readSettingsFile()
	createFolder(LOG_FOLDER)

	configureShortcuts()
	setPages()
	verifyContext()

	timeRefresh, _ := strconv.Atoi(settings.RefreshContextInformation)
	schedulerSeconds(timeRefresh, verifyContext)
}

func configureShortcuts() {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlY) {
			pagesMenu.SwitchToPage("Menu")

			app.SetFocus(listMenu)
			FooterinformationText.Clear()
		} else if event.Rune() == rune(tcell.KeyCtrlU) {
			pagesMenu.SwitchToPage(FOMR_MAINTENANCE)

			app.SetFocus(listMaintenance)
			FooterinformationText.Clear()
		} else if event.Rune() == rune(tcell.KeyCtrlQ) {
			app.Stop()
		}

		return event
	})
}

func readSettingsFile() {
	settingsFile, err := openJsonFile(SETTINGS_FILE)

	if err != nil {
		settings = AppSettings{RefreshContextInformation: "5", RefreshTablePods: "300", DefaultOutputFormatted: 0}

		writeSettingsJsonFile(settings)
	} else {
		unmarshalJson(settingsFile, &settings)
	}

	defer settingsFile.Close()
}
