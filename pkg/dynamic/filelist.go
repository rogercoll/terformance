package dynamic

import (
	"io/ioutil"
	"strconv"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"gopkg.in/yaml.v2"
)

const (
	endOfFile = "End of file"
)

type Data struct {
	Text  string `yaml:"text"`
	Color string `yaml:"color"`
}

type FileData struct {
	Content []Data `yaml:"content"`
}

// show a new file line on every update
type LineByLineTable struct {
	pos int
	ini time.Time
	FileData
	*tview.Table
}

func (i *LineByLineTable) Init() {
	return
}

func (i *LineByLineTable) Update(app *tview.Application) {
	if i.pos == 0 {
		i.ini = time.Now()
	}
	// set file line text
	text := endOfFile
	color := "white"
	if i.pos < len(i.Content) {
		text = i.Content[i.pos].Text
		color = i.Content[i.pos].Color
	}
	cell := tview.NewTableCell("").SetAlign(tview.AlignLeft).SetTextColor(tcell.GetColor(color))

	i.pos += 1
	i.InsertRow(i.pos)
	i.SetCell(i.pos, 0, tview.NewTableCell(strconv.Itoa(i.pos)).SetAlign(tview.AlignRight))
	i.SetCell(i.pos, 1, tview.NewTableCell("UVIC-UCC").SetAlign(tview.AlignRight))
	//get duration
	t2 := time.Now()
	timeDiff := t2.Sub(i.ini)
	timeOut := time.Time{}.Add(timeDiff)
	i.SetCell(i.pos, 2, tview.NewTableCell(timeOut.Format("15:04:05")).SetAlign(tview.AlignRight))

	for j, _ := range text {
		i.SetCell(i.pos, 3, cell.SetText(text[:j+1]))
		app.ForceDraw()
		time.Sleep(40 * time.Millisecond)
	}
	return
}

func NewLineByLineTable(fileName string, fileTable *tview.Table) (*LineByLineTable, error) {
	var data FileData
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &data)
	if err != nil {
		return nil, err
	}

	return &LineByLineTable{
		0,
		time.Now(),
		data,
		fileTable,
	}, nil
}
