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
	i.Update(nil)
	return
}

func (i *IncrementalText) Update(*tview.Application) {
	i.SetText(fmt.Sprintf(i.textTemplate, strconv.Itoa(i.inc)))
	i.inc += 1
	//TODO: handle ohter cases
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

type DimensionalText struct {
	inc          int
	textTemplate string
	values       []int
	*tview.TextView
}

//TODO: Implement String interface and return it in Init and Update
func (i *DimensionalText) Init() {
	i.Update(nil)
	return
}

func (i *DimensionalText) Update(*tview.Application) {
	i.SetText(fmt.Sprintf(i.textTemplate, strconv.Itoa(i.values[i.inc%len(i.values)])))
	i.inc += 1
	//TODO: handle ohter cases
	return
}

func NewDimensionalText(t string, textView *tview.TextView, values []int) *DimensionalText {
	//TODO: check if valid template
	return &DimensionalText{
		0,
		t,
		values,
		textView,
	}
}
