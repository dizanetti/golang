package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var tableContext = tview.NewTable().SetBorders(true).SetFixed(1, 1).SetSelectable(true, false)

func createTableContext() {
	tableContext.Clear()

	tableContext.SetSelectedFunc(func(row, column int) {
		currentContext := tableContext.GetCell(row, 1).Text

		if currentContext != "*" {
			context := tableContext.GetCell(row, 2).Text

			_, _, err := execPowerShellContext(context)
			if err != nil {
				FooterinformationText.SetText(err.Error()).SetTextColor(tcell.ColorRed)
			} else {
				time.Sleep(2 * time.Second)
				informationText.SetText(CURRENT_CONTEXT + context + "\n" + stringShortcuts)

				tableContext.Clear()
				configureTable(getContexts())
			}
		}
	}).SetBackgroundColor(tcell.ColorBlack)

	configureTable(getContexts())
	tableContext.ScrollToBeginning()
}

func configureTable(result []string, err error) {
	indexTable := 1
	indexHeader := 0
	indexHeaderSize := 2

	headerColumnColor := tcell.ColorWhite
	rowOKColumnColor := tcell.ColorWhite
	rowRunningColumnColor := tcell.ColorGreen

	tableContext.SetCell(0, 0, &tview.TableCell{Text: "NUMBER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableContext.SetCell(0, 1, &tview.TableCell{Text: "CURRENT", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableContext.SetCell(0, 2, &tview.TableCell{Text: "NAME", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableContext.SetCell(0, 3, &tview.TableCell{Text: "CLUSTER", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableContext.SetCell(0, 4, &tview.TableCell{Text: "AUTH. INFO", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})
	tableContext.SetCell(0, 5, &tview.TableCell{Text: "NAMESPACE", Align: tview.AlignCenter, Color: headerColumnColor, Expansion: 2, BackgroundColor: tcell.ColorGray})

	if err != nil {
		log.Fatal(err)

		return
	}

	for _, rowRaw := range result {
		infoContext := findAndDelete(splitString(rowRaw, " "), "")

		if infoContext[0] != "CURRENT" {
			for _, col := range infoContext {
				if indexHeader == 0 {
					tableContext.SetCell(indexTable, indexHeader, &tview.TableCell{Text: strconv.Itoa(indexTable), Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeader++
				}

				if indexHeader == 1 {
					if col == "*" {
						tableContext.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: rowRunningColumnColor})
					} else {
						tableContext.SetCell(indexTable, indexHeader, &tview.TableCell{Text: "", Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					}
				}

				if len(infoContext) == 3 || len(infoContext) == 4 {
					tableContext.SetCell(indexTable, indexHeaderSize, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
					indexHeaderSize++
				}
				if len(infoContext) == 5 && indexHeader != 1 {
					tableContext.SetCell(indexTable, indexHeader, &tview.TableCell{Text: col, Color: rowOKColumnColor, Expansion: 2, BackgroundColor: tcell.ColorBlack})
				}

				indexHeader++
			}
			indexTable++
			indexHeader = 0
			indexHeaderSize = 2
		}
	}
}
