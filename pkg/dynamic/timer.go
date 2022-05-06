package dynamic

import (
	"fmt"
	"time"

	"github.com/rivo/tview"
)

type Timer struct {
	textTemplate string
	ini          time.Time
	started      bool
	*tview.TextView
}

//TODO: Implement String interface and return it in Init and Update
func (i *Timer) Init() {
	i.SetText(fmt.Sprintf(i.textTemplate, "00:00:00"))
	return
}

func NewTimer(t string, textView *tview.TextView) *Timer {
	return &Timer{
		t,
		time.Now(),
		false,
		textView,
	}
}

func (i *Timer) Update(app *tview.Application) {
	if !i.started {
		i.ini = time.Now()
		//update timer every second
		go func() {
			for {
				select {
				case <-time.After(1 * time.Second):
					i.update(app)
				}
			}
		}()
		i.started = true
	}
	return
}

func (i *Timer) update(app *tview.Application) {
	//get duration
	t2 := time.Now()
	timeDiff := t2.Sub(i.ini)
	timeOut := time.Time{}.Add(timeDiff)
	i.SetText(fmt.Sprintf(i.textTemplate, timeOut.Format("15:04:05")))
	app.ForceDraw()
}
