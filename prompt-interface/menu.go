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
var informationText = tview.NewTextView().SetTextColor(tcell.ColorDarkGrey).SetTextAlign(tview.AlignCenter)

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

		pages.SwitchToPage(FORM_FILTER)
		pages.SetTitle(TITLE_FILTER)

		app.SetFocus(filterForm)
	}).
	AddItem("Pod's", "List all Pod's in context", rune(tcell.KeyCtrlP), func() {
		stringShortcuts = SHORTCUTS_PODS
		verifyContext()

		createTablePods(GET_PODS)
		tablePods.ScrollToBeginning()

		pages.SwitchToPage(FORM_POD)
		pages.SetTitle(TITLE_POD)

		app.SetFocus(tablePods)
	}).
	AddItem("Context", "Change the context", rune(tcell.KeyCtrlC), func() {
		stringShortcuts = SHORTCUTS_CONTEXT
		verifyContext()

		createTableContext()

		pages.SwitchToPage(FORM_CONTEXT)
		pages.SetTitle(TITLE_CONTEXT)

		app.SetFocus(tableContext)
	}).
	AddItem("Services", "List all Services in context", rune(tcell.KeyCtrlS), func() {
		stringShortcuts = SHORTCUTS_SERVICES
		verifyContext()

		createTableServices(GET_SERVICES)

		pages.SwitchToPage(FORM_SERVICES)
		pages.SetTitle(TITLE_SERVICE)

		app.SetFocus(tableServices)
	}).
	AddItem("Nodes", "List all Nodes in context", rune(tcell.KeyCtrlN), func() {
		stringShortcuts = SHORTCUTS_NODES
		verifyContext()

		createTableNodes(GET_NODES)

		pages.SwitchToPage(FORM_NODES)
		pages.SetTitle(TITLE_NODES)

		app.SetFocus(tableNodes)
	}).
	AddItem("Deployments", "List all Deployments in context", rune(tcell.KeyCtrlE), func() {
		stringShortcuts = SHORTCUTS_DEPLOYMENTS
		verifyContext()

		createTableDeployments(GET_DEPLOYMENTS)

		pages.SwitchToPage(FORM_DEPLOYMENTS)
		pages.SetTitle(TITLE_DEPLOYMENTS)

		app.SetFocus(tableDeployments)
	}).
	AddItem("Persistent Volumes", "List all Persistent Volumes", rune(tcell.KeyCtrlP), func() {
		stringShortcuts = SHORTCUTS_PERSISTENT_VOLUMES
		verifyContext()

		createTablePersistentVolumes(GET_PERSISTENT_VOLUMES + " " + GET_PERSISTENT_VOLUMES_ARGS)

		pages.SwitchToPage(FORM_PERSISTENT_VOLUMES)
		pages.SetTitle(TITLE_PERSISTENT_VOLUMES)

		app.SetFocus(tablePersistentVolumes)
	}).
	AddItem("Config Maps", "List all Config Maps", rune(tcell.KeyCtrlG), func() {
		stringShortcuts = SHORTCUTS_CONFIG_MAPS
		verifyContext()

		createTableConfigMaps(GET_CONFIG_MAPS)

		pages.SwitchToPage(FORM_CONFIG_MAPS)
		pages.SetTitle(TITLE_CONFIG_MAPS)

		app.SetFocus(tableConfigMaps)
	}).
	AddItem("Maintenance", "Functions to POD maintenance", rune(tcell.KeyCtrlM), func() {
		stringShortcuts = SHORTCUTS_MAINTENANCE
		verifyContext()

		pagesMenu.SwitchToPage(FOMR_MAINTENANCE)
		pagesMenu.SetTitle(TITLE_MAINTENANCE)

		app.SetFocus(listMaintenance)
	}).
	AddItem("App Settings", "Settings", rune(tcell.KeyCtrlS), func() {
		stringShortcuts = SHORTCUTS_SETTINGS
		verifyContext()

		setSettingsForm()

		pages.SwitchToPage(FORM_APP_SETTINGS)
		pages.SetTitle(TITLE_APP_SETTINGS)

		app.SetFocus(settingsForm)
	}).
	AddItem("Help", "Informations", rune(tcell.KeyCtrlH), func() {
		stringShortcuts = ""
		verifyContext()

		pages.SetTitle("")
		pages.SwitchToPage(FORM_HELP)
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
	pagesMenu.AddPage(FOMR_MAINTENANCE, listMaintenance, true, true)
	pagesMenu.AddPage("Menu", listMenu, true, true)

	listMenu.SetBorder(true).SetTitle("Menu")
	listMaintenance.SetBorder(true).SetTitle(TITLE_MAINTENANCE)

	pages.AddPage(FORM_FILTER, filterForm, true, true).SetBorder(true)
	pages.AddPage("CopyLogsFromPodForm", copyLogsFromPodForm, true, true).SetBorder(true)

	pages.AddPage(FORM_APP_SETTINGS, settingsForm, true, true).SetBorder(true)

	pages.AddPage("Teste C", textC, true, true).SetBorder(true)
	pages.AddPage("DescribePod", describePod, true, true).SetBorder(true)
	pages.AddPage("LoadConfiguration", loadConfiguration, true, true).SetBorder(true)

	pages.AddPage(FORM_CONTEXT, tableContext, true, true).SetBorder(true)
	pages.AddPage(FORM_SERVICES, tableServices, true, true).SetBorder(true)
	pages.AddPage(FORM_NODES, tableNodes, true, true).SetBorder(true)
	pages.AddPage(FORM_DEPLOYMENTS, tableDeployments, true, true).SetBorder(true)
	pages.AddPage(FORM_CONFIG_MAPS, tableConfigMaps, true, true).SetBorder(true)
	pages.AddPage(FORM_PERSISTENT_VOLUMES, tablePersistentVolumes, true, true).SetBorder(true)
	pages.AddPage(FORM_POD, tablePods, true, true).SetBorder(true)

	pages.AddPage("ModalSettingsButtonOK", createModalSettingsButtonOK(), true, true)
	pages.AddPage(FORM_HELP, welcomeText, true, true).SetBorder(true)

	infoPages.AddPage("InformationText", informationText, true, true).SetBorder(true).SetTitle(TITLE_CONTEXT)

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

			pages.SwitchToPage(FORM_POD)
			pages.SetTitle(TITLE_POD)

			app.SetFocus(tablePods)
		} else if valueType == 1 { //Service
			stringShortcuts = SHORTCUTS_CONTEXT
			verifyContext()

			createTableServices(GET_SERVICES, valueName)

			pages.SwitchToPage(FORM_SERVICES)
			pages.SetTitle(TITLE_SERVICE)

			app.SetFocus(tableServices)
		} else if valueType == 2 { //Deployments
			stringShortcuts = SHORTCUTS_DEPLOYMENTS
			verifyContext()

			createTableDeployments(GET_DEPLOYMENTS, valueName)

			pages.SwitchToPage(FORM_DEPLOYMENTS)
			pages.SetTitle(TITLE_DEPLOYMENTS)

			app.SetFocus(tableDeployments)
		} else if valueType == 3 { //ConfigMaps
			stringShortcuts = SHORTCUTS_CONFIG_MAPS
			verifyContext()

			createTableConfigMaps(GET_CONFIG_MAPS, valueName)

			pages.SwitchToPage(FORM_CONFIG_MAPS)
			pages.SetTitle(TITLE_CONFIG_MAPS)

			app.SetFocus(tableConfigMaps)
		}
	})

	filterForm.AddButton("Cancel", func() {

		createTablePods(GET_PODS)
		pages.SwitchToPage(FORM_POD)
		app.SetFocus(tablePods)
	})
}
