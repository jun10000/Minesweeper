package main

import (
	"jun10000.github.io/minesweeper/container2"
	"jun10000.github.io/minesweeper/widget2"

	"fmt"
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
			window.SetContent(newGameLayout(int(width), int(height), int(bombs)))
			window.Resize(fyne.NewSize(0, 0))
		}),
	)
}

func newGameLayout(width int, height int, bombs int) *fyne.Container {
	return widget2.NewMSTable(width, height, bombs,
		func() {
			window.SetContent(newClearLayout(width, height, bombs))
			window.Resize(fyne.NewSize(240, 0))
		},
		func() {
			window.SetContent(newGameOverLayout(width, height, bombs))
			window.Resize(fyne.NewSize(240, 0))
		})
}

func newClearLayout(width int, height int, bombs int) *fyne.Container {
	return container.NewVBox(
		widget.NewLabelWithStyle("Clear!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("(%dx%d, %dBombs)", width, height, bombs), fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("Elapsed Time: ---", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewButton("Replay?", func() {
			window.SetContent(newTitleLayout())
			window.Resize(fyne.NewSize(640, 0))
		}),
	)
}

func newGameOverLayout(width int, height int, bombs int) *fyne.Container {
	return container.NewVBox(
		widget.NewLabelWithStyle("GameOver...", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(fmt.Sprintf("(%dx%d, %dBombs)", width, height, bombs), fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewLabelWithStyle("Elapsed Time: ---", fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewButton("Replay?", func() {
			window.SetContent(newTitleLayout())
			window.Resize(fyne.NewSize(640, 0))
		}),
	)
}

func main() {
	application = app.New()
	window = application.NewWindow(TITLE)
	window.SetContent(newTitleLayout())
	window.Resize(fyne.NewSize(640, 0))
	window.ShowAndRun()
}
