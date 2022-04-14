package terformance

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type TUI struct {
	App  *tview.Application
	Grid *tview.Grid

	Persons DynamicText
	Road    DynamicText
	Objects DynamicText
	Music   *tview.TextView

	Schema *tview.List
}

func NewTUI() *TUI {
	t := TUI{}
	t.App = tview.NewApplication()

	t.Persons = NewIncrementalText("[yellow::bl]Persons:[-:-:-]"+" [ %s ]", tview.NewTextView().SetDynamicColors(true))
	t.Road = NewIncrementalText("Road: %s%%", tview.NewTextView().SetTextColor(tcell.ColorAqua))
	t.Objects = NewIncrementalText("Objects: %s", tview.NewTextView().SetTextColor(tcell.ColorAqua))
	t.Music = tview.NewTextView().SetTextColor(tcell.ColorAqua)

	t.Schema = tview.NewList().ShowSecondaryText(true).SetSecondaryTextColor(tcell.ColorRed).SetSelectedStyle(tcell.Style{}.Foreground(tcell.ColorBlack).Background(tcell.ColorWhite))

	// Configure appearance
	t.Persons.Init()
	t.Road.Init()
	t.Objects.Init()
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
		SetRows(1).
		SetBorders(true).
		AddItem(header, 0, 0, 2, 2, 0, 0, true).
		AddItem(messages, 3, 0, 22, 2, 0, 0, false)

	//t.setBackground(tcell.ColorBlack)
	t.setupKeyboard()

	return &t
}

// Start starts terminal user interface application.
func (tui *TUI) Start() error {
	return tui.App.SetRoot(tui.Grid, true).EnableMouse(false).Run()
}

func (tui *TUI) LoadData() {
	tui.Road.Update()
	tui.Persons.Update()
}
