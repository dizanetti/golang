package main

import (
	"github.com/rivo/tview"
)

type AppSettings struct {
	RefreshTablePods          string `json:"refresh_table_pods"`
	RefreshContextInformation string `json:"refresh_context_information"`
	LogFolder                 string `json:"log_folder"`
	DefaultOutputFormatted    int    `json:"default_output_formatted"`
}

var settingsForm = tview.NewForm()

func setSettingsForm() {
	var timeContextValue string = settings.RefreshContextInformation
	var timePodsValue string = settings.RefreshTablePods
	var logFolderPath string = settings.LogFolder
	var defaultOutputFormatted int = settings.DefaultOutputFormatted

	settingsForm.Clear(true)

	settingsForm.AddInputField("Time to refresh Context in Area Information(in seconds)", settings.RefreshContextInformation, 25, nil, func(timeContext string) {
		timeContextValue = timeContext
	})

	settingsForm.AddInputField("Time to refresh Pod's(in seconds)", settings.RefreshTablePods, 25, nil, func(timePods string) {
		timePodsValue = timePods
	})

	settingsForm.AddInputField("Path to Log folder in Pod", settings.LogFolder, 25, nil, func(path string) {
		logFolderPath = path
	})

	settingsForm.AddDropDown("Default output formatted API object", []string{"YAML", "JSON"}, settings.DefaultOutputFormatted, func(option string, optionIndex int) {
		defaultOutputFormatted = optionIndex
	})

	settingsForm.AddButton("Ok", func() {
		settings.RefreshContextInformation = timeContextValue
		settings.RefreshTablePods = timePodsValue
		settings.DefaultOutputFormatted = defaultOutputFormatted
		settings.LogFolder = logFolderPath

		app.SetFocus(modalAppSettingsConfirm)
		pages.SwitchToPage("ModalSettingsButtonOK")
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
