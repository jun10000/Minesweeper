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
