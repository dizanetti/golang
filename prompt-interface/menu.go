package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var stringShortcuts string = ""

var pages = tview.NewPages()

var textC = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("Work in Progress...")
var welcomeText = tview.NewTextView().SetTextColor(tcell.ColorYellow).SetText(openTextFile(WELCOME_BANNER)).SetTextAlign(tview.AlignCenter)
var describePod = tview.NewTextView().SetTextColor(tcell.ColorYellow).SetScrollable(true)
var loadConfiguration = tview.NewTextView().SetTextColor(tcell.ColorYellow).SetScrollable(true)

var infoPages = tview.NewPages()
var informationText = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetTextAlign(tview.AlignCenter)

var filterForm = tview.NewForm()

var modalAppSettingsConfirm = tview.NewModal()

var list = tview.NewList().
	AddItem("Pod's", "List all Pod's in context", rune(tcell.KeyCtrlP), func() {
		stringShortcuts = SHORTCUTS_PODS
		verifyContext()

		createTablePods(GET_PODS)
		tablePods.ScrollToBeginning()

		pages.SwitchToPage("TablesPods")
		pages.SetTitle("Pod's")

		app.SetFocus(tablePods)
	}).
	AddItem("Context", "Change the context", rune(tcell.KeyCtrlC), func() {
		stringShortcuts = SHORTCUTS_CONTEXT
		verifyContext()

		createTableContext()

		pages.SwitchToPage("TablesContext")
		pages.SetTitle("Context")

		app.SetFocus(tableContext)
	}).
	AddItem("Services", "List all Services in context", rune(tcell.KeyCtrlS), func() {
		stringShortcuts = SHORTCUTS_SERVICES
		verifyContext()

		createTableServices(GET_SERVICES)

		pages.SwitchToPage("TablesServices")
		pages.SetTitle("Services")

		app.SetFocus(tableServices)
	}).
	AddItem("Nodes", "List all Nodes in context", rune(tcell.KeyCtrlN), func() {
		stringShortcuts = SHORTCUTS_NODES
		verifyContext()

		createTableNodes(GET_NODES)

		pages.SwitchToPage("TablesNodes")
		pages.SetTitle("Nodes")

		app.SetFocus(tableNodes)
	}).
	AddItem("Deployments", "List all Deployments in context", rune(tcell.KeyCtrlE), func() {
		stringShortcuts = SHORTCUTS_DEPLOYMENTS
		verifyContext()

		createTableDeployments(GET_DEPLOYMENTS)

		pages.SwitchToPage("TablesDeployments")
		pages.SetTitle("Deployments")

		app.SetFocus(tableDeployments)
	}).
	AddItem("Persistent Volumes", "List all Persistent Volumes", rune(tcell.KeyCtrlP), func() {
		stringShortcuts = SHORTCUTS_PERSISTENT_VOLUMES
		verifyContext()

		createTablePersistentVolumes(GET_PERSISTENT_VOLUMES + " " + GET_PERSISTENT_VOLUMES_ARGS)

		pages.SwitchToPage("TablesPersistentVolumes")
		pages.SetTitle("Persistent Volumes")

		app.SetFocus(tablePersistentVolumes)
	}).
	AddItem("Filter", "Filter a list of Pod's/Services", rune(tcell.KeyCtrlF), func() {
		stringShortcuts = SHORTCUTS_FILTER
		verifyContext()

		setFilterForm()

		pages.SwitchToPage("filterForm")
		pages.SetTitle("Filter")

		app.SetFocus(filterForm)
	}).
	AddItem("Maintenance", "Functions to POD maintenance", rune(tcell.KeyCtrlM), func() {
		pages.SwitchToPage("Teste C")
	}).
	AddItem("App Settings", "Settings", rune(tcell.KeyCtrlS), func() {
		stringShortcuts = SHORTCUTS_SETTINGS
		verifyContext()

		setSettingsForm()

		pages.SwitchToPage("AppSettingsForm")
		pages.SetTitle("Settings")

		app.SetFocus(settingsForm)
	}).
	AddItem("Help", "Informations", rune(tcell.KeyCtrlH), func() {
		stringShortcuts = ""
		verifyContext()

		pages.SetTitle("")
		pages.SwitchToPage("Help")
	}).
	AddItem("Quit", "Press to exit", rune(tcell.KeyCtrlQ), func() {
		app.Stop()
	})

func setPages() {
	list.SetBorder(true).SetTitle("Menu")

	pages.AddPage("filterForm", filterForm, true, true).SetBorder(true)
	pages.AddPage("AppSettingsForm", settingsForm, true, true).SetBorder(true)

	pages.AddPage("Teste C", textC, true, true).SetBorder(true)
	pages.AddPage("DescribePod", describePod, true, true).SetBorder(true)
	pages.AddPage("LoadConfiguration", loadConfiguration, true, true).SetBorder(true)

	pages.AddPage("TablesContext", tableContext, true, true).SetBorder(true)
	pages.AddPage("TablesServices", tableServices, true, true).SetBorder(true)
	pages.AddPage("TablesNodes", tableNodes, true, true).SetBorder(true)
	pages.AddPage("TablesDeployments", tableDeployments, true, true).SetBorder(true)
	pages.AddPage("TablesPersistentVolumes", tablePersistentVolumes, true, true).SetBorder(true)
	pages.AddPage("TablesPods", tablePods, true, true).SetBorder(true)

	pages.AddPage("ModalSettingsButtonOK", createModalSettingsButtonOK(), true, true)
	pages.AddPage("Help", welcomeText, true, true).SetBorder(true)

	infoPages.AddPage("InformationText", informationText, true, true).SetBorder(true).SetTitle("Information")

	describePodShortcuts()
}

func setFilterForm() {
	var valueName string
	var valueType int

	filterForm.Clear(true)
	filterForm.AddInputField("Name", "", 25, nil, func(name string) {
		valueName = name
	})

	filterForm.AddDropDown("Type", []string{"Pod", "Service"}, 0, func(option string, optionIndex int) {
		valueType = optionIndex
	})

	filterForm.AddButton("Find", func() {
		if valueType == 0 {
			stringShortcuts = SHORTCUTS_PODS
			verifyContext()

			createTablePods(GET_PODS, valueName)

			pages.SwitchToPage("TablesPods")
			pages.SetTitle("Pod's")

			app.SetFocus(tablePods)
		} else if valueType == 1 {
			stringShortcuts = SHORTCUTS_CONTEXT
			verifyContext()

			createTableServices(GET_SERVICES, valueName)

			pages.SwitchToPage("TablesServices")
			pages.SetTitle("Services")

			app.SetFocus(tableServices)
		}
	})

	filterForm.AddButton("Cancel", func() {

		createTablePods(GET_PODS)
		pages.SwitchToPage("TablesPods")
		app.SetFocus(tablePods)
	})
}

func describePodShortcuts() {
	describePod.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlM) {
			pages.SwitchToPage("TablesPods")

			stringShortcuts = SHORTCUTS_PODS
			verifyContext()

			app.SetFocus(tablePods)
		}
		return event
	})
}
