package main

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tableDeployments = tview.NewTable().SetBorders(true).SetFixed(1, 1).SetSelectable(true, false)

func createTableDeployments(commands ...string) {
	tableDeployments.Clear()
	tableDeployments.SetBackgroundColor(tcell.ColorBlack)

	configureTabletableDeployments(execute(commands...))
	tableDeployments.ScrollToBeginning()
}

func configureTabletableDeployments(result []string, err error) {
	indexTable := 1
	indexHeader := 0

	headerColumnColor := tcell.ColorWhite
	rowOKColumnColor := tcell.ColorWhite

	tableDeployments.SetCell(0, 0, &tview.TableCell{Text: "NUMBER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableDeployments.SetCell(0, 1, &tview.TableCell{Text: "NAME", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableDeployments.SetCell(0, 2, &tview.TableCell{Text: "READY", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableDeployments.SetCell(0, 3, &tview.TableCell{Text: "UP-TO-DATE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableDeployments.SetCell(0, 4, &tview.TableCell{Text: "AVAILABLE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableDeployments.SetCell(0, 5, &tview.TableCell{Text: "AGE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})

	if err != nil {
		log.Fatal(err)

		return
	}

	for _, rowRaw := range result {
		infoServices := findAndDelete(splitString(rowRaw, " "), "")

		if infoServices[0] != "NAME" {
			for _, col := range infoServices {
				if indexHeader == 0 {
					tableDeployments.SetCell(indexTable, indexHeader, &tview.TableCell{Text: strconv.Itoa(indexTable), Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeader++
				}
				tableDeployments.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})

				indexHeader++
			}
			indexTable++
			indexHeader = 0
		}
	}
}
