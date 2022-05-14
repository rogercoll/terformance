package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rogercoll/terformance/internal/config"
	"github.com/rogercoll/terformance/pkg/dynamic"
)

type TUI struct {
	App  *tview.Application
	Grid *tview.Grid

	Persons    dynamic.DynamicItem
	University dynamic.DynamicItem

	Light dynamic.DynamicItem
	// static text
	File *tview.TextView

	// static text
	Date *tview.TextView
	Time dynamic.DynamicItem

	Schema dynamic.DynamicItem
}

func newTable(columnsTitles []string) *tview.Table {
	t := tview.NewTable().SetFixed(1, len(columnsTitles))
	color := tcell.ColorBlack
	bckColor := tcell.ColorWhite
	for c, title := range columnsTitles {
		if c == 3 {
			t.SetCell(0, c,
				tview.NewTableCell(title).
					SetTextColor(color).
					SetBackgroundColor(bckColor).
					SetExpansion(18).
					SetAlign(tview.AlignLeft))
		} else {
			t.SetCell(0, c,
				tview.NewTableCell(title).
					SetTextColor(color).
					SetBackgroundColor(bckColor).
					SetMaxWidth(14).
					SetAlign(tview.AlignCenter))
		}
	}
	return t
}

func NewTUI(cfg config.AppConfig) (*TUI, error) {
	t := TUI{}
	t.App = tview.NewApplication()

	t.Persons = dynamic.NewIncrementalText("[#3DFFF3::b]Persones:[-:-:-]"+" [ %s ]", tview.NewTextView().SetDynamicColors(true))
	t.University = dynamic.NewDimensionalText("[#3DFFF3::b]Universitat:[-:-:-]\t%s", tview.NewTextView().SetDynamicColors(true), []int{1717, 1977, 1997, 2022})
	t.Light = dynamic.NewDimensionalText("[#3DFFF3::b]Llums encesos:[-:-:-]\t[%s/25]", tview.NewTextView().SetDynamicColors(true), []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25})
	t.File = tview.NewTextView().SetDynamicColors(true).SetText("C:/[#FF1B1B::b]25AnysUVIC-UCC[-:-:-]/Pre-lu-di")
	t.Date = tview.NewTextView().SetText("20/05/2022")
	t.Time = dynamic.NewTimer("%s", tview.NewTextView().SetDynamicColors(true))

	var err error
	t.Schema, err = dynamic.NewLineByLineTable(cfg.FileName, newTable([]string{"TASCA", "USUARI", "TEMPS+", "COMANDAMENT"}))
	if err != nil {
		return nil, err
	}

	// Configure appearance
	t.Persons.Init()
	t.University.Init()
	t.Light.Init()
	t.Time.Init()

	// Input handlers

	// Layout
	header := tview.NewGrid().SetRows(0, 0).
		AddItem(t.Persons, 0, 0, 1, 1, 0, 0, false).
		AddItem(t.University, 1, 0, 1, 1, 0, 0, false).
		AddItem(t.Light, 0, 1, 1, 1, 0, 0, false).
		AddItem(t.File, 1, 1, 1, 1, 0, 0, false).
		AddItem(t.Date, 0, 2, 1, 1, 0, 0, false).
		AddItem(t.Time, 1, 2, 1, 1, 0, 0, false)
	messages := tview.NewGrid().SetRows(0).
		AddItem(t.Schema, 0, 0, 1, 1, 0, 0, false)
	t.Grid = tview.NewGrid().
		SetRows(1).
		SetBorders(false).
		AddItem(header, 0, 0, 2, 3, 0, 0, true).
		AddItem(messages, 3, 0, 22, 3, 0, 0, false)

	//t.setBackground(tcell.ColorBlack)
	t.setupKeyboard()

	return &t, nil
}

// Start starts terminal user interface application.
func (tui *TUI) Start() error {
	return tui.App.SetRoot(tui.Grid, true).EnableMouse(false).Run()
}
