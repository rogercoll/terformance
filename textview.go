package terformance

import "github.com/rivo/tview"

type DynamicText interface {
	Init() string
	Update() string
}

type IncrementalText struct {
	inc  int
	text string
	*tview.TextView
}

func (i *IncrementalText) Init() string {
	return i.text
}

func (i *IncrementalText) Update() string {
	return i.text
}
