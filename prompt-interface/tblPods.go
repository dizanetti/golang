package main

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tablePods = tview.NewTable().SetBorders(true).SetFixed(1, 1).SetSelectable(true, false)

func createTablePods(commands ...string) {
	tablePods.Clear()

	tablePods.SetSelectedFunc(func(row, column int) {
		podName := tablePods.GetCell(row, 1).Text

		_, _, err := execPowerShell(podName)
		if err != nil {
			informationText.SetText(err.Error()).SetTextColor(tcell.ColorRed)
		}
	}).SetBackgroundColor(tcell.ColorBlack)

	indexTable := 1
	indexHeader := 0

	headerColumnColor := tcell.ColorWhite

	rowOKColumn := tcell.ColorWhite
	rowRunningColumn := tcell.ColorGreen
	rowNOKColumn := tcell.ColorRed

	result, err := execute(commands...)

	tablePods.SetCell(0, 0, &tview.TableCell{Text: "NUMBER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tablePods.SetCell(0, 1, &tview.TableCell{Text: "NAME", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tablePods.SetCell(0, 2, &tview.TableCell{Text: "READY", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tablePods.SetCell(0, 3, &tview.TableCell{Text: "STATUS", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tablePods.SetCell(0, 4, &tview.TableCell{Text: "RESTARTS", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tablePods.SetCell(0, 5, &tview.TableCell{Text: "AGE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})

	if err != nil {
		log.Fatal(err)

		return
	}

	for _, rowRaw := range result {
		infoPod := findAndDelete(splitString(rowRaw, " "), "")

		if infoPod[0] != "NAME" {
			for _, col := range infoPod {
				if indexHeader == 0 {
					tablePods.SetCell(indexTable, indexHeader, &tview.TableCell{Text: strconv.Itoa(indexTable), Color: rowOKColumn, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeader++
				}

				if indexHeader == 3 {
					if col == "Running" {
						tablePods.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumn, Expansion: 2, BackgroundColor: rowRunningColumn})
					} else {
						tablePods.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumn, Expansion: 2, BackgroundColor: rowNOKColumn})
					}
				} else {
					tablePods.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumn, Expansion: 2, BackgroundColor: tcell.ColorBlack})
				}

				indexHeader++
			}
			indexTable++
			indexHeader = 0
		}
	}
}
