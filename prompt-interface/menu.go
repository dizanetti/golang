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
	AddItem("List all Pod's", "List all Pod's in context", 'l', func() {
		createTablePods(GET_PODS)

		pages.SwitchToPage("TablesPods")
		pages.SetTitle("Pod's")

		app.SetFocus(tablePods)
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
	AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

func setPages() {
	list.SetBorder(true).SetTitle("Menu")

	pages.AddPage("filterForm", filterForm, true, true).SetBorder(true)
	pages.AddPage("Teste C", textC, true, true).SetBorder(true)
	pages.AddPage("TablesContext", tableContext, true, true).SetBorder(true)
	pages.AddPage("TablesPods", tablePods, true, true).SetBorder(true)
	pages.AddPage("Help", welcomeText, true, true).SetBorder(true)

	infoPages.AddPage("InformationText", informationText, true, true).SetBorder(true).SetTitle("Information")
}

func setFilterForm() {
	var value string

	filterForm.Clear(true)
	filterForm.AddInputField("Pod Name", "", 25, nil, func(podName string) {
		value = podName
	})

	filterForm.AddButton("Find", func() {
		createTablePods(GET_PODS, value)
		pages.SwitchToPage("TablesPods")
		app.SetFocus(tablePods)
	})

	filterForm.AddButton("Cancel", func() {

		createTablePods(GET_PODS)
		pages.SwitchToPage("TablesPods")
		app.SetFocus(tablePods)
	})
}
