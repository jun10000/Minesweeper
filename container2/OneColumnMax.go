package container2

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type OneColumnMax struct {
	MaxIndex int
}

func NewOneColumnMax(index int, objects ...fyne.CanvasObject) *fyne.Container {
	return container.New(&OneColumnMax{MaxIndex: index}, objects...)
}

func (c *OneColumnMax) MinSize(objects []fyne.CanvasObject) fyne.Size {
	width, height := float32(0), float32(0)

	for _, o := range objects {
		s := o.MinSize()
		width += s.Width
		if s.Height > height {
			height = s.Height
		}
	}

	return fyne.NewSize(width, height)
}

func (c *OneColumnMax) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	x := float32(0)

	for i, o := range objects {
		// Deside size
		s := o.MinSize()
		if i == c.MaxIndex {
			s.Width = containerSize.Width - c.MinSize(objects).Width + s.Width
		}
		o.Resize(s)

		// Deside position
		y := (containerSize.Height - s.Height) / 2
		o.Move(fyne.NewPos(x, y))

		// For next object
		x += s.Width
	}
}
