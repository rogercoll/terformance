package dynamic

import "github.com/rivo/tview"

type DynamicItem interface {
	Init()
	Update(*tview.Application)
	tview.Primitive
}
