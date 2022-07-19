package main

import (
	"github.com/rivo/tview"
)

type AppSettings struct {
	RefreshTablePods          string `json:"refresh_table_pods"`
	RefreshContextInformation string `json:"refresh_context_information"`
}

var settingsForm = tview.NewForm()

func setSettingsForm() {
	var timeContextValue string
	var timePodsValue string

	settingsForm.Clear(true)

	settingsForm.AddInputField("Time to refresh Context in Area Information(in seconds)", settings.RefreshContextInformation, 25, nil, func(timeContext string) {
		timeContextValue = timeContext
	})

	settingsForm.AddInputField("Time to refresh Pod's(in seconds)", settings.RefreshTablePods, 25, nil, func(timePods string) {
		timePodsValue = timePods
	})

	settingsForm.AddButton("Ok", func() {
		app.SetFocus(modalAppSettingsConfirm)
		pages.SwitchToPage("ModalSettingsButtonOK")

		if timeContextValue != "" {
			settings.RefreshContextInformation = timeContextValue
		}

		if timePodsValue != "" {
			settings.RefreshTablePods = timePodsValue
		}
	})

	settingsForm.AddButton("Cancel", func() {
		settingsForm.Clear(true)

		setSettingsForm()
		app.SetFocus(settingsForm)
	})
}

func createModalSettingsButtonOK() *tview.Modal {
	modalAppSettingsConfirm = tview.NewModal().
		SetText("Do you want to save the settings? \n \nYou need to restart the application.").
		AddButtons([]string{"Confirm", "Cancel"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "Confirm" {
				writeSettingsJsonFile(settings)

				pages.SwitchToPage("AppSettingsForm")
				pages.SetTitle("Settings")

				app.SetFocus(settingsForm)
			} else {
				pages.SwitchToPage("AppSettingsForm")
				pages.SetTitle("Settings")

				app.SetFocus(settingsForm)
			}
		})

	return modalAppSettingsConfirm
}
