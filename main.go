package main

import (
	"jun10000.github.io/minesweeper/container2"
	"jun10000.github.io/minesweeper/widget2"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const TITLE = "Minesweeper"

var (
	application fyne.App
	window fyne.Window
)

func newTitleLayout() *fyne.Container {
	width := float64(13)
	height := float64(5)
	bombs := float64(10)

	width_data := binding.BindFloat(&width)
	height_data := binding.BindFloat(&height)
	bombs_data := binding.BindFloat(&bombs)

	return container.NewVBox(
		widget.NewLabelWithStyle(TITLE, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container2.NewOneColumnMax(1,
			widget.NewLabel("Width"),
			widget.NewSliderWithData(2, 30, width_data),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(width_data, "%.0f")),
		),
		container2.NewOneColumnMax(1,
			widget.NewLabel("Height"),
			widget.NewSliderWithData(2, 30, height_data),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(height_data, "%.0f")),
		),
		container2.NewOneColumnMax(1,
			widget.NewLabel("Bombs"),
			widget.NewSliderWithData(1, 300, bombs_data),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(bombs_data, "%.0f")),
		),
		widget.NewButton("START", func() {
			window.Resize(fyne.NewSize(0, 0))
			window.SetContent(newGameLayout(int(width), int(height), int(bombs)))
		}),
	)
}

func newGameLayout(width int, height int, bombs int) *fyne.Container {
	return widget2.NewMSTable(width, height, bombs)
}

func main() {
	application = app.New()
	window = application.NewWindow(TITLE)
	window.Resize(fyne.NewSize(640, 0))
	window.SetContent(newTitleLayout())
	window.ShowAndRun()
}
