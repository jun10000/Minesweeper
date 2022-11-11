package widget2

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

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
	HasBomb bool
	NearBombs int
}

func NewMSCell() *MSCell {
	c := &MSCell {}
	c.ExtendBaseWidget(c)
	return c
}

func (c *MSCell) Tapped(e *fyne.PointEvent) {
	if c.IsOpened || c.MarkState != MSCellMarkStatesNone {
		return
	}

	println("左クリック押下")
	if c.HasBomb {
		c.Button.Text = "x"
	} else {
		if c.NearBombs > 0 {
			c.Button.Text = fmt.Sprint(c.NearBombs)
		}
	}

	c.IsOpened = true
	c.Refresh()
}

func (c *MSCell) TappedSecondary(e *fyne.PointEvent) {
	if c.IsOpened {
		return
	}

	println("右クリック押下")
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
