package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var pages = tview.NewPages()

var textC = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetText("WIP")
var welcomeText = tview.NewTextView().SetTextColor(tcell.ColorDarkGrey).SetText(openTextFile(WELCOME_BANNER)).SetTextAlign(tview.AlignCenter)

var infoPages = tview.NewPages()
var informationText = tview.NewTextView().SetTextColor(tcell.ColorGreen).SetTextAlign(tview.AlignCenter)

var filterForm = tview.NewForm()

var list = tview.NewList().
	AddItem("Pod's", "List all Pod's in context", 'l', func() {
		createTablePods(GET_PODS)

		pages.SwitchToPage("TablesPods")
		pages.SetTitle("Pod's")

		app.SetFocus(tablePods)
	}).
	AddItem("Services", "List all Services in context", 's', func() {
		createTableServices(GET_SERVICES)

		pages.SwitchToPage("TablesServices")
		pages.SetTitle("Services")

		app.SetFocus(tableServices)
	}).
	AddItem("Filter", "Filter a list of Pod's", 'f', func() {
		setFilterForm()

		pages.SwitchToPage("filterForm")
		pages.SetTitle("Filter")

		app.SetFocus(filterForm)
	}).
	AddItem("Maintenance", "Functions to help with POD maintenance", 'm', func() {
		pages.SwitchToPage("Teste C")
	}).
	AddItem("Context", "Change the context", 'c', func() {
		createTableContext()

		pages.SwitchToPage("TablesContext")
		pages.SetTitle("Context")

		app.SetFocus(tableContext)
	}).
	AddItem("Help", "Informations", 'i', func() {
		pages.SetTitle("")
		pages.SwitchToPage("Help")
	}).
	AddItem("Quit", "Press to exit", 'h', func() {
		app.Stop()
	})

func setPages() {
	list.SetBorder(true).SetTitle("Menu")

	pages.AddPage("filterForm", filterForm, true, true).SetBorder(true)
	pages.AddPage("Teste C", textC, true, true).SetBorder(true)
	pages.AddPage("TablesContext", tableContext, true, true).SetBorder(true)
	pages.AddPage("TablesServices", tableServices, true, true).SetBorder(true)
	pages.AddPage("TablesPods", tablePods, true, true).SetBorder(true)
	pages.AddPage("Help", welcomeText, true, true).SetBorder(true)

	infoPages.AddPage("InformationText", informationText, true, true).SetBorder(true).SetTitle("Information")
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
			createTablePods(GET_PODS, valueName)

			pages.SwitchToPage("TablesPods")
			pages.SetTitle("Pod's")

			app.SetFocus(tablePods)
		} else if valueType == 1 {
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
