package widget2

import (
	"jun10000.github.io/minesweeper/resource"
	"jun10000.github.io/minesweeper/utility"

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
	Seed int64
	OnClear func(time.Duration, utility.Position)
	OnGameOver func(time.Duration, utility.Position)
	NonOpenedCells int
	Status MSTableStates
	Cells *[]fyne.CanvasObject
	InitTime time.Time
	FirstCell *MSCell
}

func NewMSTable(width int, height int, bombs int, seed int64, onClear func(time.Duration, utility.Position), onGameOver func(time.Duration, utility.Position)) *fyne.Container {
	t := &MSTable{
		Width: width,
		Height: height,
		Bombs: bombs,
		Seed: seed,
		OnClear: onClear,
		OnGameOver: onGameOver,
		NonOpenedCells: width * height,
		Status: MSTableStatesNonInit,
	}

	c := container.New(t)
	t.Cells = &c.Objects

	for i := 0; i < (width * height); i++ {
		c.Add(NewMSCell(t, i))
	}

	return c
}

func (t *MSTable) Init(exceptCell *MSCell) {
	bombArray := utility.GetRandBinaryArrayWithSeed(t.Width * t.Height - 1, t.Bombs, t.Seed)
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
	t.FirstCell = exceptCell
}

//++++++++++++++++++++++++++++++
// MSTable Container Methods
//++++++++++++++++++++++++++++++

func (t *MSTable) MinSize(objects []fyne.CanvasObject) fyne.Size {
	if len(objects) == 0 {
		return fyne.NewSize(0, 0)
	}

	cellSize := objects[0].MinSize()
	w := cellSize.Width * float32(t.Width)
	h := cellSize.Height * float32(t.Height)

	return fyne.NewSize(w, h)
}

func (t *MSTable) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) == 0 {
		return
	}

	cellSize := objects[0].MinSize()

	for _, o := range objects {
		p := o.(*MSCell).GetPosition()
		x := float32(p.X) * cellSize.Width
		y := float32(p.Y) * cellSize.Height

		o.Resize(cellSize)
		o.Move(fyne.NewPos(x, y))
	}
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
		t.OnGameOver(time.Now().Sub(t.InitTime), t.FirstCell.GetPosition())
		return
	}

	if t.NonOpenedCells <= t.Bombs {
		t.Status = MSTableStatesClear
		t.OnClear(time.Now().Sub(t.InitTime), t.FirstCell.GetPosition())
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
	widget.Icon
	
	Parent *MSTable
	Index int
	HasBomb bool

	IsOpened bool
	MarkState MSCellMarkStates
}

func NewMSCell(parent *MSTable, index int) *MSCell {
	c := &MSCell {Parent: parent, Index: index}
	c.ExtendBaseWidget(c)
	c.Refresh()
	return c
}

func (c *MSCell) GetPosition() utility.Position {
	return utility.NewPosition(c.Index % c.Parent.Width, c.Index / c.Parent.Width)
}

func (c *MSCell) GetIndex(pos utility.Position) int {
	return c.Parent.Width * pos.Y + pos.X
}

func (c *MSCell) GetNearCells() []*MSCell {
	pos := c.GetPosition()
	poslist := [8]utility.Position {
		utility.NewPosition(pos.X - 1, pos.Y - 1),
		utility.NewPosition(pos.X    , pos.Y - 1),
		utility.NewPosition(pos.X + 1, pos.Y - 1),
		utility.NewPosition(pos.X - 1, pos.Y    ),
		utility.NewPosition(pos.X + 1, pos.Y    ),
		utility.NewPosition(pos.X - 1, pos.Y + 1),
		utility.NewPosition(pos.X    , pos.Y + 1),
		utility.NewPosition(pos.X + 1, pos.Y + 1),
	}
	ret := make([]*MSCell, 0, 8)

	for _, p := range poslist {
		if p.X < 0 || p.X >= c.Parent.Width || p.Y < 0 || p.Y >= c.Parent.Height {
			continue
		}

		cell := (*c.Parent.Cells)[c.GetIndex(p)].(*MSCell)
		ret = append(ret, cell)
	}

	return ret
}

func (c *MSCell) GetNearBombs() int {
	count := 0

	for _, cell := range c.GetNearCells() {
		if cell.HasBomb {
			count++
		}
	}

	return count
}

func (c *MSCell) Open() {
	if c.IsOpened || c.Parent.Status == MSTableStatesClear {
		return
	}

	c.Parent.OnCellOpen(c)
	c.IsOpened = true
	c.Refresh()

	// Recursive calls
	if !c.HasBomb && c.GetNearBombs() == 0 {
		for _, cell := range c.GetNearCells() {
			cell.Open()
		}
	}
}

func (c *MSCell) SwitchMarkState() {
	switch c.MarkState {
	case MSCellMarkStatesNone:
		c.MarkState = MSCellMarkStatesBomb
	case MSCellMarkStatesBomb:
		c.MarkState = MSCellMarkStatesQuestion
	case MSCellMarkStatesQuestion:
		c.MarkState = MSCellMarkStatesNone
	}

	c.Refresh()
}

//++++++++++++++++++++++++++++++
// MSCell Widget Methods
//++++++++++++++++++++++++++++++

func (c *MSCell) MinSize() fyne.Size {
	return fyne.NewSize(32, 32)
}

func (c *MSCell) Tapped(e *fyne.PointEvent) {
	if c.MarkState != MSCellMarkStatesNone {
		return
	}

	c.Open()
}

func (c *MSCell) TappedSecondary(e *fyne.PointEvent) {
	if c.IsOpened || c.Parent.Status == MSTableStatesClear {
		return
	}

	c.SwitchMarkState()
}

func (c *MSCell) Refresh() {
	var r *fyne.StaticResource = nil

	if c.IsOpened {
		if c.HasBomb {
			r = resource.OpenbombPng
		} else {
			switch c.GetNearBombs() {
			case 0:
				r = resource.OpennonePng
			case 1:
				r = resource.Opennear1Png
			case 2:
				r = resource.Opennear2Png
			case 3:
				r = resource.Opennear3Png
			case 4:
				r = resource.Opennear4Png
			case 5:
				r = resource.Opennear5Png
			case 6:
				r = resource.Opennear6Png
			case 7:
				r = resource.Opennear7Png
			case 8:
				r = resource.Opennear8Png
			}
		}
	} else {
		switch c.MarkState {
		case MSCellMarkStatesNone:
			r = resource.ClosenonePng
		case MSCellMarkStatesBomb:
			r = resource.ClosebombPng
		case MSCellMarkStatesQuestion:
			r = resource.ClosequestionPng
		}
	}

	c.Icon.SetResource(r)
}
