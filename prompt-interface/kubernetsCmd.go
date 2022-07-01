package main

import "github.com/gdamore/tcell/v2"

func verifyContext() {
	contextK8, errContextK8 := execCmd("kubectl config current-context").String()
	if errContextK8 != nil {
		informationText.SetText("Error message: " + errContextK8.Error()).SetTextColor(tcell.ColorRed)
	} else {
		informationText.SetText(CURRENT_CONTEXT + contextK8 + stringShortcuts).SetTextColor(tcell.ColorGreen)
	}
}

func getContexts() ([]string, error) {
	return execute("kubectl config get-contexts")
}
