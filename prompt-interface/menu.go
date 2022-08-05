package main

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"golang.org/x/exp/slices"
)

var stringShortcuts string = ""

var pages = tview.NewPages()
var pagesMenu = tview.NewPages()

var textC = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("Work in Progress...")
var welcomeText = tview.NewTextView().SetTextColor(tcell.ColorYellow).SetText(openTextFile(WELCOME_BANNER)).SetTextAlign(tview.AlignCenter)
var describePod = tview.NewTextView().SetTextColor(tcell.ColorYellow).SetScrollable(true)
var loadConfiguration = tview.NewTextView().SetTextColor(tcell.ColorYellow).SetScrollable(true)

var infoPages = tview.NewPages()
var informationText = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetTextAlign(tview.AlignCenter)

var FooterPages = tview.NewPages()
var FooterinformationText = tview.NewTextView().SetTextColor(tcell.ColorRed).SetTextAlign(tview.AlignCenter)

var filterForm = tview.NewForm()
var copyLogsFromPodForm = tview.NewForm()

var modalAppSettingsConfirm = tview.NewModal()

var listMenu = tview.NewList().
	AddItem("Filter", "Filter a list of Pod's/Services", rune(tcell.KeyCtrlF), func() {
		stringShortcuts = SHORTCUTS_FILTER
		verifyContext()

		setFilterForm()

		pages.SwitchToPage("filterForm")
		pages.SetTitle("Filter")

		app.SetFocus(filterForm)
	}).
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
	AddItem("Config Maps", "List all Config Maps", rune(tcell.KeyCtrlG), func() {
		stringShortcuts = SHORTCUTS_CONFIG_MAPS
		verifyContext()

		createTableConfigMaps(GET_CONFIG_MAPS)

		pages.SwitchToPage("TablesConfigMaps")
		pages.SetTitle("Config Maps")

		app.SetFocus(tableConfigMaps)
	}).
	AddItem("Maintenance", "Functions to POD maintenance", rune(tcell.KeyCtrlM), func() {
		stringShortcuts = SHORTCUTS_MAINTENANCE
		verifyContext()

		pagesMenu.SwitchToPage("Maintenance")
		pagesMenu.SetTitle("Maintenance")

		app.SetFocus(listMaintenance)
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

var listMaintenance = tview.NewList().
	AddItem("Logs", "Copy Logs from Pod's to local", rune(tcell.KeyCtrlL), func() {
		stringShortcuts = SHORTCUTS_COPY_LOGS
		verifyContext()

		setCopyLogsFromPodForm()

		pages.SwitchToPage("CopyLogsFromPodForm")
		pages.SetTitle("Copy Log's")

		app.SetFocus(copyLogsFromPodForm)
	})

func setPages() {
	pagesMenu.AddPage("Maintenance", listMaintenance, true, true)
	pagesMenu.AddPage("Menu", listMenu, true, true)

	listMenu.SetBorder(true).SetTitle("Menu")
	listMaintenance.SetBorder(true).SetTitle("Maintenance")

	pages.AddPage("filterForm", filterForm, true, true).SetBorder(true)
	pages.AddPage("CopyLogsFromPodForm", copyLogsFromPodForm, true, true).SetBorder(true)

	pages.AddPage("AppSettingsForm", settingsForm, true, true).SetBorder(true)

	pages.AddPage("Teste C", textC, true, true).SetBorder(true)
	pages.AddPage("DescribePod", describePod, true, true).SetBorder(true)
	pages.AddPage("LoadConfiguration", loadConfiguration, true, true).SetBorder(true)

	pages.AddPage("TablesContext", tableContext, true, true).SetBorder(true)
	pages.AddPage("TablesServices", tableServices, true, true).SetBorder(true)
	pages.AddPage("TablesNodes", tableNodes, true, true).SetBorder(true)
	pages.AddPage("TablesDeployments", tableDeployments, true, true).SetBorder(true)
	pages.AddPage("TablesConfigMaps", tableConfigMaps, true, true).SetBorder(true)
	pages.AddPage("TablesPersistentVolumes", tablePersistentVolumes, true, true).SetBorder(true)
	pages.AddPage("TablesPods", tablePods, true, true).SetBorder(true)

	pages.AddPage("ModalSettingsButtonOK", createModalSettingsButtonOK(), true, true)
	pages.AddPage("Help", welcomeText, true, true).SetBorder(true)

	infoPages.AddPage("InformationText", informationText, true, true).SetBorder(true).SetTitle("Context")

	FooterPages.AddPage("FooterinformationText", FooterinformationText, true, true).SetBorder(true).SetTitle("Information")
}

func setCopyLogsFromPodForm() {
	var podName string

	copyLogsFromPodForm.Clear(true)

	result, err := execCmcReturnSliceAndColumn(GET_PODS, 1)
	if err != nil {
		FooterinformationText.SetText(err.Error())
	} else {
		result = slices.Delete(result, 0, 1)

		copyLogsFromPodForm.AddDropDown("Select Pod", result, 0, func(option string, optionIndex int) {
			podName = option
		})
	}

	copyLogsFromPodForm.AddButton("Copy", func() {
		folder := LOG_FOLDER + "/" + podName

		errRemoveFolder := removeFolder(folder)
		if errRemoveFolder == nil {
			errFolder := createFolder(folder)

			if errFolder != nil {
				FooterinformationText.SetText(errFolder.Error())
			} else {
				pathPod := strings.Replace(settings.LogFolder, "{POD_NAME}", podName, -1)

				_, err1, err2 := execPowerShellCopyFiles(podName, folder, pathPod)
				if err2 != nil {
					FooterinformationText.SetText(err1 + " | " + err2.Error())
				}
			}
		}
	})

	copyLogsFromPodForm.AddButton("Cancel", func() {
	})
}

func setFilterForm() {
	var valueName string
	var valueType int

	filterForm.Clear(true)
	filterForm.AddInputField("Name", "", 25, nil, func(name string) {
		valueName = name
	})

	filterForm.AddDropDown("Type", []string{"Pod", "Service", "Deployments", "ConfigMaps"}, 0, func(option string, optionIndex int) {
		valueType = optionIndex
	})

	filterForm.AddButton("Find", func() {
		if valueType == 0 { //Pod
			stringShortcuts = SHORTCUTS_PODS
			verifyContext()

			createTablePods(GET_PODS, valueName)

			pages.SwitchToPage("TablesPods")
			pages.SetTitle("Pod's")

			app.SetFocus(tablePods)
		} else if valueType == 1 { //Service
			stringShortcuts = SHORTCUTS_CONTEXT
			verifyContext()

			createTableServices(GET_SERVICES, valueName)

			pages.SwitchToPage("TablesServices")
			pages.SetTitle("Services")

			app.SetFocus(tableServices)
		} else if valueType == 2 { //Deployments
			stringShortcuts = SHORTCUTS_DEPLOYMENTS
			verifyContext()

			createTableDeployments(GET_DEPLOYMENTS, valueName)

			pages.SwitchToPage("TablesDeployments")
			pages.SetTitle("Deployments")

			app.SetFocus(tableDeployments)
		} else if valueType == 3 { //ConfigMaps
			stringShortcuts = SHORTCUTS_CONFIG_MAPS
			verifyContext()

			createTableConfigMaps(GET_CONFIG_MAPS, valueName)

			pages.SwitchToPage("TablesConfigMaps")
			pages.SetTitle("Config Maps")

			app.SetFocus(tableConfigMaps)
		}
	})

	filterForm.AddButton("Cancel", func() {

		createTablePods(GET_PODS)
		pages.SwitchToPage("TablesPods")
		app.SetFocus(tablePods)
	})
}
