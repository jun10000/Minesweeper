package widget2

import (
	"jun10000.github.io/minesweeper/utility"

	"fmt"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

//++++++++++++++++++++++++++++++
// MSTable
//++++++++++++++++++++++++++++++

type MSTableStates int
const (
	MSTableStatesNonInit MSTableStates = iota
	MSTableStatesInit
	MSTableStatesClear
)

type MSTable struct {
	Width int
	Height int
	Bombs int
	OnClear func(time.Duration)
	OnGameOver func(time.Duration)
	Cells *[]fyne.CanvasObject
	NonOpenedCells int
	Status MSTableStates
	InitTime time.Time
}

func NewMSTable(width int, height int, bombs int, onClear func(time.Duration), onGameOver func(time.Duration)) *fyne.Container {
	c := container.NewGridWithColumns(width)
	t := &MSTable{
		Width: width,
		Height: height,
		Bombs: bombs,
		OnClear: onClear,
		OnGameOver: onGameOver,
		Cells: &c.Objects,
		NonOpenedCells: width * height,
		Status: MSTableStatesNonInit,
	}

	for i := 0; i < (width * height); i++ {
		c.Add(NewMSCell(t, i))
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

	t.Status = MSTableStatesInit
	t.InitTime = time.Now()
}

//++++++++++++++++++++++++++++++
// MSTable Callback Methods
//++++++++++++++++++++++++++++++

func (t *MSTable) OnCellOpen(c *MSCell) {
	if t.Status == MSTableStatesNonInit {
		t.Init(c)
	}

	t.NonOpenedCells--

	if c.HasBomb {
		t.Status = MSTableStatesClear
		t.OnGameOver(time.Now().Sub(t.InitTime))
		return
	}

	if t.NonOpenedCells <= t.Bombs {
		t.Status = MSTableStatesClear
		t.OnClear(time.Now().Sub(t.InitTime))
		return
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
	
	Parent *MSTable
	Index int
	HasBomb bool

	IsOpened bool
	MarkState MSCellMarkStates
}

func NewMSCell(parent *MSTable, index int) *MSCell {
	c := &MSCell {Parent: parent, Index: index}
	c.ExtendBaseWidget(c)
	c.Text = " "
	return c
}

func (c *MSCell) GetPosition() utility.Position {
	return utility.Position {X: c.Index % c.Parent.Width, Y: c.Index / c.Parent.Width}
}

func (c *MSCell) GetIndex(pos utility.Position) int {
	return c.Parent.Width * pos.Y + pos.X
}

func (c *MSCell) GetNearBombs() int {
	count := 0
	pos := c.GetPosition()
	aroundPositions := [8]utility.Position {
		{X: pos.X - 1, Y: pos.Y - 1},
		{X: pos.X    , Y: pos.Y - 1},
		{X: pos.X + 1, Y: pos.Y - 1},
		{X: pos.X - 1, Y: pos.Y    },
		{X: pos.X + 1, Y: pos.Y    },
		{X: pos.X - 1, Y: pos.Y + 1},
		{X: pos.X    , Y: pos.Y + 1},
		{X: pos.X + 1, Y: pos.Y + 1},
	}

	for _, p := range aroundPositions {
		if p.X < 0 || p.X >= c.Parent.Width || p.Y < 0 || p.Y >= c.Parent.Height {
			continue
		}

		cell := (*c.Parent.Cells)[c.GetIndex(p)].(*MSCell)
		if cell.HasBomb {
			count++
		}
	}

	return count
}

//++++++++++++++++++++++++++++++
// MSCell Override Methods
//++++++++++++++++++++++++++++++

func (c *MSCell) Tapped(e *fyne.PointEvent) {
	if c.IsOpened || c.MarkState != MSCellMarkStatesNone || c.Parent.Status == MSTableStatesClear {
		return
	}

	c.Parent.OnCellOpen(c)
	
	if c.HasBomb {
		c.Button.Text = "x"
	} else {
		c.Button.Text = fmt.Sprint(c.GetNearBombs())
	}

	c.IsOpened = true
	c.Refresh()
}

func (c *MSCell) TappedSecondary(e *fyne.PointEvent) {
	if c.IsOpened || c.Parent.Status == MSTableStatesClear {
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
