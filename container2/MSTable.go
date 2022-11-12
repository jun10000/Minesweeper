package container2

import (
	"jun10000.github.io/minesweeper/widget2"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type MSTable struct {
	Width int
	Height int
	Bombs int
	Initialized bool
}

func NewMSTable(width int, height int, bombs int) *fyne.Container {
	c := container.NewGridWithColumns(width)
	for i := 0; i < (width * height); i++ {
		c.Add(widget2.NewMSCell())
	}

	return container.New(&MSTable{Width: width, Height: height, Bombs: bombs}, c)
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
