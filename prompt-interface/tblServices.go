package main

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tableServices = tview.NewTable().SetBorders(true).SetFixed(1, 1).SetSelectable(true, false)

func createTableServices(commands ...string) {
	tableServices.Clear()
	tableServices.SetBackgroundColor(tcell.ColorBlack).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlL) {
			row, _ := tableServices.GetSelection()

			podName := tableServices.GetCell(row, 1).Text

			load(podName, "services", tableServices, "TablesServices", SHORTCUTS_SERVICES)
		}

		return event
	})

	configureTableServices(execute(commands...))
	tableServices.ScrollToBeginning()
}

func configureTableServices(result []string, err error) {
	indexTable := 1
	indexHeader := 0

	headerColumnColor := tcell.ColorWhite
	rowOKColumnColor := tcell.ColorWhite

	tableServices.SetCell(0, 0, &tview.TableCell{Text: "NUMBER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableServices.SetCell(0, 1, &tview.TableCell{Text: "NAME", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableServices.SetCell(0, 2, &tview.TableCell{Text: "TYPE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableServices.SetCell(0, 3, &tview.TableCell{Text: "CLUSTER-IP", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableServices.SetCell(0, 4, &tview.TableCell{Text: "EXTERNAL-IP", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableServices.SetCell(0, 5, &tview.TableCell{Text: "PORT(S)", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableServices.SetCell(0, 6, &tview.TableCell{Text: "AGE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})

	if err != nil {
		log.Fatal(err)

		return
	}

	for _, rowRaw := range result {
		infoServices := findAndDelete(splitString(rowRaw, " "), "")

		if infoServices[0] != "NAME" {
			for _, col := range infoServices {
				if indexHeader == 0 {
					tableServices.SetCell(indexTable, indexHeader, &tview.TableCell{Text: strconv.Itoa(indexTable), Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeader++
				}
				tableServices.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})

				indexHeader++
			}
			indexTable++
			indexHeader = 0
		}
	}
}
