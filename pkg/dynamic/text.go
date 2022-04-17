package dynamic

import (
	"fmt"
	"strconv"

	"github.com/rivo/tview"
)

type IncrementalText struct {
	inc          int
	textTemplate string
	*tview.TextView
}

//TODO: Implement String interface and return it in Init and Update
func (i *IncrementalText) Init() {
	i.SetText(fmt.Sprintf(i.textTemplate, strconv.Itoa(i.inc)))
	return
}

func (i *IncrementalText) Update() {
	i.inc += 1
	i.SetText(fmt.Sprintf(i.textTemplate, strconv.Itoa(i.inc)))
	return
}

func NewIncrementalText(t string, textView *tview.TextView) *IncrementalText {
	//TODO: check if valid template
	return &IncrementalText{
		0,
		t,
		textView,
	}
}
