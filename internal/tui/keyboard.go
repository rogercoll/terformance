package tui

import "github.com/gdamore/tcell/v2"

func (tui *TUI) setupKeyboard() {
	tui.App.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlA:
			tui.Persons.Update(tui.App)
		case tcell.KeyCtrlB:
			tui.University.Update(tui.App)
		case tcell.KeyCtrlD:
			tui.Light.Update(tui.App)
		case tcell.KeyCtrlE:
			tui.Schema.Update(tui.App)
			tui.Time.Update(tui.App)
		}
		return event
	})
}
