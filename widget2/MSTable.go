package widget2

import (
	"jun10000.github.io/minesweeper/utility"

	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//++++++++++++++++++++++++++++++
// MSTable
//++++++++++++++++++++++++++++++

type MSTable struct {
	Width int
	Height int
	Bombs int
	Cells *[]fyne.CanvasObject
	IsInit bool
}

func NewMSTable(width int, height int, bombs int) *fyne.Container {
	c := container.NewGridWithColumns(width)
	t := &MSTable{Width: width, Height: height, Bombs: bombs, Cells: &c.Objects}

	for i := 0; i < (width * height); i++ {
		c.Add(NewMSCell(t))
	}

	return container.New(t, c)
}

func (t *MSTable) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) != 1 {
		return fyne.NewSize(0, 0)
	}

	return objects[0].MinSize()
}

func (t *MSTable) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) != 1 {
		return
	}

	o := objects[0]
	o.Resize(containerSize)
	o.Move(fyne.NewPos(0, 0))
}

func (t *MSTable) Init(exceptCell *MSCell) {
	bombArray := utility.GetRandBinaryArray(t.Width * t.Height - 1, t.Bombs)
	currentBombIndex := 0
	for _, c := range *t.Cells {
		if c == exceptCell {
			continue
		}
		cell := c.(*MSCell)
		cell.HasBomb = bombArray[currentBombIndex]
		currentBombIndex++
	}

	// cell.NearBombs = ???
	// ToDo: Sum around 8 cell's bombs
	//     Get around 8 cell's index
	//         Set index into all cells on initializing cells

	t.IsInit = true
}

//++++++++++++++++++++++++++++++
// MSTable Callback Methods
//++++++++++++++++++++++++++++++

func (t *MSTable) OnCellOpen(c *MSCell) {
	if !t.IsInit {
		t.Init(c)
	}
}

//++++++++++++++++++++++++++++++
// MSCell
//++++++++++++++++++++++++++++++

type MSCellMarkStates int
const (
	MSCellMarkStatesNone MSCellMarkStates = iota
	MSCellMarkStatesBomb
	MSCellMarkStatesQuestion
)

type MSCell struct {
	widget.Button

	IsOpened bool
	MarkState MSCellMarkStates

	PosX int
	PosY int
	Parent *MSTable
	HasBomb bool
	NearBombs int
}

func NewMSCell(parent *MSTable) *MSCell {
	c := &MSCell {Parent: parent}
	c.ExtendBaseWidget(c)
	c.Text = " "
	return c
}

func (c *MSCell) Tapped(e *fyne.PointEvent) {
	if c.IsOpened || c.MarkState != MSCellMarkStatesNone {
		return
	}

	if c.HasBomb {
		c.Button.Text = "x"
	} else {
		c.Button.Text = fmt.Sprint(c.NearBombs)
	}

	c.IsOpened = true
	c.Parent.OnCellOpen(c)
	c.Refresh()
}

func (c *MSCell) TappedSecondary(e *fyne.PointEvent) {
	if c.IsOpened {
		return
	}

	switch c.MarkState {
	case MSCellMarkStatesNone:
		c.MarkState = MSCellMarkStatesBomb
		c.Button.Text = "B"
	case MSCellMarkStatesBomb:
		c.MarkState = MSCellMarkStatesQuestion
		c.Button.Text = "?"
	case MSCellMarkStatesQuestion:
		c.MarkState = MSCellMarkStatesNone
		c.Button.Text = " "
	}

	c.Refresh()
}
