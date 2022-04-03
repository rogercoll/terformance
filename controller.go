package terformance

type Controller struct {
	t *TUI
}

func LoadConfig() *Controller {
	return &Controller{NewTUI()}
}

func (c *Controller) Run() {
	c.t.Start()
}
