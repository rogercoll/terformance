package terformance

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App  *tview.Application
	Grid *tview.Grid

	Persons *tview.TextView
	Road    *tview.TextView
	Objects *tview.TextView
	Music   *tview.TextView

	Schema *tview.List
}

func (t *TUI) setBackground(color tcell.Color) {
	t.Persons.Box = tview.NewBox().SetBackgroundColor(color)
	t.Road.Box = tview.NewBox().SetBackgroundColor(color)
	t.Objects.Box = tview.NewBox().SetBackgroundColor(color)
	t.Music.Box = tview.NewBox().SetBackgroundColor(tcell.ColorYellow)

	t.Grid.Box = tview.NewBox().SetBackgroundColor(tcell.ColorBlack)
}

func NewTUI() *TUI {
	t := TUI{}
	t.App = tview.NewApplication()

	t.Persons = tview.NewTextView().SetTextColor(tcell.ColorAqua)
	t.Road = tview.NewTextView().SetTextColor(tcell.ColorAqua)
	t.Objects = tview.NewTextView().SetTextColor(tcell.ColorAqua)
	t.Music = tview.NewTextView().SetTextColor(tcell.ColorAqua)

	t.Schema = tview.NewList().ShowSecondaryText(true).SetSecondaryTextColor(tcell.ColorRed).SetSelectedStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite))

	// Configure appearance
	t.Persons.SetText("Persons")
	t.Road.SetText("Road")
	t.Objects.SetText("Objects")
	t.Music.SetText("Music")

	// Input handlers

	// Layout
	header := tview.NewGrid().SetRows(0, 0).
		AddItem(t.Persons, 0, 0, 1, 1, 0, 0, false).
		AddItem(t.Road, 1, 0, 1, 1, 0, 0, false).
		AddItem(t.Objects, 0, 1, 1, 1, 0, 0, false).
		AddItem(t.Music, 1, 1, 1, 1, 0, 0, false)
	messages := tview.NewGrid().SetRows(0).
		AddItem(t.Schema, 0, 0, 1, 1, 0, 0, false)
	t.Grid = tview.NewGrid().
		SetRows(0, 25).
		SetBorders(true).
		AddItem(header, 0, 0, 1, 1, 0, 0, true).
		AddItem(messages, 1, 0, 2, 1, 0, 0, false)

	t.setBackground(tcell.ColorBlack)
	t.setupKeyboard()

	return &t
}

// Start starts terminal user interface application.
func (tui *TUI) Start() error {
	return tui.App.SetRoot(tui.Grid, true).EnableMouse(false).Run()
}

func (tui *TUI) LoadData() {
	tui.Road.SetText("Roadddddddd")
}
