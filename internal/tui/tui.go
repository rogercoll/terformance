package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rogercoll/terformance/pkg/dynamic"
)

type TUI struct {
	App  *tview.Application
	Grid *tview.Grid

	Persons dynamic.DynamicItem
	Road    dynamic.DynamicItem
	Objects dynamic.DynamicItem
	Music   dynamic.DynamicItem

	Schema dynamic.DynamicItem
}

func newTable(columnsTitles []string) *tview.Table {
	t := tview.NewTable().SetFixed(1, len(columnsTitles))
	color := tcell.ColorBlack
	bckColor := tcell.ColorRed
	for c, title := range columnsTitles {
		expansion := 1
		if c == 3 {
			expansion = 8
		}
		t.SetCell(0, c,
			tview.NewTableCell(title).
				SetTextColor(color).
				SetBackgroundColor(bckColor).
				SetExpansion(expansion).
				SetAlign(tview.AlignCenter))
	}
	return t
}

func NewTUI() *TUI {
	t := TUI{}
	t.App = tview.NewApplication()

	t.Persons = dynamic.NewIncrementalText("[yellow::b]Persons:[-:-:-]"+" [ %s ]", tview.NewTextView().SetDynamicColors(true))
	t.Road = dynamic.NewIncrementalText("Road: %s%%", tview.NewTextView().SetTextColor(tcell.ColorAqua))
	t.Objects = dynamic.NewIncrementalText("Objects: %s", tview.NewTextView().SetTextColor(tcell.ColorAqua))
	t.Music = dynamic.NewIncrementalText("Music: %s", tview.NewTextView().SetTextColor(tcell.ColorAqua))

	t.Schema = dynamic.NewLineByLineTable(newTable([]string{"TASCA", "USUARI", "TEMPS+", "COMMAND"}))

	// Configure appearance
	t.Persons.Init()
	t.Road.Init()
	t.Objects.Init()
	t.Music.Init()

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
		SetBorders(false).
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
