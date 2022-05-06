package dynamic

import (
	"bufio"
	"os"
	"strconv"
	"time"

	"github.com/rivo/tview"
)

const (
	endOfFile = "End of file"
)

// show a new file line on every update
type LineByLineTable struct {
	pos  int
	ini  time.Time
	data []string
	*tview.Table
}

func (i *LineByLineTable) Init() {
	return
}

func (i *LineByLineTable) Update(app *tview.Application) {
	if i.pos == 0 {
		i.ini = time.Now()
	}
	i.pos += 1
	i.InsertRow(i.pos)
	i.SetCell(i.pos, 0, tview.NewTableCell(strconv.Itoa(i.pos)).SetAlign(tview.AlignRight))
	i.SetCell(i.pos, 1, tview.NewTableCell("UVIC-UCC").SetAlign(tview.AlignRight))
	//get duration
	t2 := time.Now()
	timeDiff := t2.Sub(i.ini)
	timeOut := time.Time{}.Add(timeDiff)
	i.SetCell(i.pos, 2, tview.NewTableCell(timeOut.Format("15:04:05")).SetAlign(tview.AlignRight))
	// set file line text
	text := endOfFile
	if i.pos < len(i.data) {
		text = i.data[i.pos]
	}
	cell := tview.NewTableCell("").SetAlign(tview.AlignCenter)
	for j, _ := range text {
		i.SetCell(i.pos, 3, cell.SetText(text[:j]))
		app.ForceDraw()
		time.Sleep(40 * time.Millisecond)
	}
	return
}

func NewLineByLineTable(fileName string, fileTable *tview.Table) (*LineByLineTable, error) {
	var fileLines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &LineByLineTable{
		0,
		time.Now(),
		fileLines,
		fileTable,
	}, nil
}
