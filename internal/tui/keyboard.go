package tui

import "github.com/gdamore/tcell/v2"

func (tui *TUI) setupKeyboard() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.Road.Update()
		case tcell.KeyCtrlB:
			tui.Persons.Update()
		case tcell.KeyCtrlD:
			tui.Schema.Update()
		}
		return event
	})
}