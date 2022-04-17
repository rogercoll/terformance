package controller

import "github.com/rogercoll/terformance/internal/tui"

type Controller struct {
	t *tui.TUI
}

func LoadConfig() *Controller {
	//TODO: pass config to tui
	return &Controller{tui.NewTUI()}
}

func (c *Controller) Run() {
	c.t.Start()
}
