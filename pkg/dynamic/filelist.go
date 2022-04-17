package dynamic

import "github.com/rivo/tview"

// show a new file line on every update
type LineByLineTable struct {
	pos int
	*tview.Table
}

func (i *LineByLineTable) Init() {
	return
}

func (i *LineByLineTable) Update() {
	i.pos += 1
	i.InsertRow(i.pos)
	i.SetCell(i.pos, 2, tview.NewTableCell("hello"))
	return
}

func NewLineByLineTable(fileTable *tview.Table) *LineByLineTable {
	//TODO: check if valid template
	return &LineByLineTable{
		0,
		fileTable,
	}
}
