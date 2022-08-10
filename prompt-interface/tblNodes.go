package main

import (
	"log"
	"strconv"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tableNodes = tview.NewTable().SetBorders(true).SetFixed(1, 1).SetSelectable(true, false)

func createTableNodes(commands ...string) {
	tableNodes.Clear()
	tableNodes.SetBackgroundColor(tcell.ColorBlack).SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == rune(tcell.KeyCtrlL) {
			row, _ := tableNodes.GetSelection()

			podName := tableNodes.GetCell(row, 1).Text

			load(podName, "nodes", tableNodes, FORM_NODES, SHORTCUTS_NODES, TITLE_NODES)
		} else if event.Rune() == rune(tcell.KeyCtrlI) {
			row, _ := tableNodes.GetSelection()

			nodeName := tableNodes.GetCell(row, 1).Text

			describeResult, errDescribe := executeKubectlCore("describe", "nodes", nodeName)
			if errDescribe != "" {
				FooterinformationText.SetText(errDescribe).SetTextColor(tcell.ColorRed)
			} else {
				describe(describeResult, tableNodes, FORM_NODES, SHORTCUTS_DESCRIBE, TITLE_NODES)
			}
		}

		return event
	})

	configureTabletableNodes(execute(commands...))
	tableNodes.ScrollToBeginning()
}

func configureTabletableNodes(result []string, err error) {
	indexTable := 1
	indexHeader := 0

	headerColumnColor := tcell.ColorWhite
	rowOKColumnColor := tcell.ColorWhite

	tableNodes.SetCell(0, 0, &tview.TableCell{Text: "NUMBER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableNodes.SetCell(0, 1, &tview.TableCell{Text: "NAME", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableNodes.SetCell(0, 2, &tview.TableCell{Text: "STATUS", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableNodes.SetCell(0, 3, &tview.TableCell{Text: "ROLES", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableNodes.SetCell(0, 4, &tview.TableCell{Text: "AGE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableNodes.SetCell(0, 5, &tview.TableCell{Text: "VERSION", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})

	if err != nil {
		log.Fatal(err)

		return
	}

	for _, rowRaw := range result {
		infoServices := findAndDelete(splitString(rowRaw, " "), "")

		if infoServices[0] != "NAME" {
			for _, col := range infoServices {
				if indexHeader == 0 {
					tableNodes.SetCell(indexTable, indexHeader, &tview.TableCell{Text: strconv.Itoa(indexTable), Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeader++
				}
				tableNodes.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})

				indexHeader++
			}
			indexTable++
			indexHeader = 0
		}
	}
}
