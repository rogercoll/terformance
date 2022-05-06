package tui

import "github.com/gdamore/tcell/v2"

func (tui *TUI) setupKeyboard() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.Time.Update(tui.App)
		case tcell.KeyCtrlB:
			tui.Persons.Update(tui.App)
		case tcell.KeyCtrlD:
			tui.Schema.Update(tui.App)
			tui.Time.Update(tui.App)
		}
		return event
	})
}
