package main

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tableConfigMaps = tview.NewTable().SetBorders(true).SetFixed(1, 1).SetSelectable(true, false)

func createTableConfigMaps(commands ...string) {
	tableConfigMaps.Clear()
	tableConfigMaps.SetBackgroundColor(tcell.ColorBlack).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlL) {
			row, _ := tableConfigMaps.GetSelection()

			podName := tableConfigMaps.GetCell(row, 1).Text

			load(podName, "configmap", tableConfigMaps, "TablesConfigMaps", SHORTCUTS_CONFIG_MAPS)
		} else if event.Rune() == rune(tcell.KeyCtrlI) {
			row, _ := tableConfigMaps.GetSelection()

			podName := tableConfigMaps.GetCell(row, 1).Text

			describeResult, errDescribe := executeKubectlCore("describe", "configmaps", podName)
			if errDescribe != "" {
				informationText.SetText(errDescribe).SetTextColor(tcell.ColorRed)
			} else {
				describe(describeResult, tableConfigMaps, "TablesConfigMaps", SHORTCUTS_CONFIG_MAPS)
			}
		}

		return event
	})

	configureTableConfigMaps(execute(commands...))
	tableConfigMaps.ScrollToBeginning()
}

func configureTableConfigMaps(result []string, err error) {
	indexTable := 1
	indexHeader := 0

	headerColumnColor := tcell.ColorWhite
	rowOKColumnColor := tcell.ColorWhite

	tableConfigMaps.SetCell(0, 0, &tview.TableCell{Text: "NUMBER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableConfigMaps.SetCell(0, 1, &tview.TableCell{Text: "NAME", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableConfigMaps.SetCell(0, 2, &tview.TableCell{Text: "DATA", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableConfigMaps.SetCell(0, 3, &tview.TableCell{Text: "AGE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})

	if err != nil {
		log.Fatal(err)

		return
	}

	for _, rowRaw := range result {
		infoServices := findAndDelete(splitString(rowRaw, " "), "")

		if infoServices[0] != "NAME" {
			for _, col := range infoServices {
				if indexHeader == 0 {
					tableConfigMaps.SetCell(indexTable, indexHeader, &tview.TableCell{Text: strconv.Itoa(indexTable), Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeader++
				}
				tableConfigMaps.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})

				indexHeader++
			}
			indexTable++
			indexHeader = 0
		}
	}
}
