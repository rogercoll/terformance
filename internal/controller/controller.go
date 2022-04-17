package controller

import (
	"github.com/rogercoll/terformance/internal/config"
	"github.com/rogercoll/terformance/internal/tui"
)

type Controller struct {
	t *tui.TUI
}

func LoadConfig(fileName string) (*Controller, error) {
	//TODO: pass config to tui
	cfg := config.AppConfig{fileName}
	t, err := tui.NewTUI(cfg)
	if err != nil {
		return nil, err
	}
	return &Controller{t}, nil
}

func (c *Controller) Run() {
	c.t.Start()
}
