package main

import (
	"jun10000.github.io/minesweeper/container2"
	"jun10000.github.io/minesweeper/utility"
	"jun10000.github.io/minesweeper/widget2"

	"fmt"
	"time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const TITLE = "Minesweeper"

var (
	application fyne.App
	window_main fyne.Window
	window_result fyne.Window

	width = float64(13)
	height = float64(5)
	bombs = float64(10)
	seed = float64(time.Now().UnixNano())
)

func newTitleLayout() *fyne.Container {
	width_data := binding.BindFloat(&width)
	height_data := binding.BindFloat(&height)
	bombs_data := binding.BindFloat(&bombs)
	seed_data := binding.BindFloat(&seed)

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
		container.NewHBox(
			widget.NewLabel("Seed"),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(seed_data, "%.0f")),
			widget.NewButton("Refresh", func() {
				seed_data.Set(float64(time.Now().UnixNano()))
			}),
		),
		widget.NewButton("START", func() {
			window_main.SetContent(newGameLayout())
			window_main.Resize(fyne.NewSize(0, 0))
		}),
	)
}

func newGameLayout() *fyne.Container {
	return widget2.NewMSTable(int(width), int(height), int(bombs), int64(seed),
		func(elapsedTime time.Duration, firstPos utility.Position) {
			window_result = application.NewWindow("Result")
			window_result.SetContent(newClearLayout(elapsedTime, firstPos))
			window_result.Resize(fyne.NewSize(300, 0))
			window_result.Show()
		},
		func(elapsedTime time.Duration, firstPos utility.Position) {
			window_result = application.NewWindow("Result")
			window_result.SetContent(newGameOverLayout(elapsedTime, firstPos))
			window_result.Resize(fyne.NewSize(300, 0))
			window_result.Show()
		})
}

func newClearLayout(elapsedTime time.Duration, firstPos utility.Position) *fyne.Container {
	et_h, et_m, et_s := utility.SplitDuration(elapsedTime)
	return container.NewVBox(
		widget.NewLabelWithStyle("Clear!", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(fmt.Sprintf(
			"Size: %.0fx%.0f\n" +
			"Bombs: %.0f\n" +
			"Seed: %.0f\n" +
			"First Cell: (%d,%d)\n" +
			"Elapsed Time: %d:%02d:%02d",
			width, height, bombs, seed, firstPos.X, firstPos.Y, et_h, et_m, et_s), fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewButton("Replay?", func() {
			window_main.SetContent(newTitleLayout())
			window_main.Resize(fyne.NewSize(640, 0))
			window_result.Close()
		}),
	)
}

func newGameOverLayout(elapsedTime time.Duration, firstPos utility.Position) *fyne.Container {
	et_h, et_m, et_s := utility.SplitDuration(elapsedTime)
	return container.NewVBox(
		widget.NewLabelWithStyle("GameOver...", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabelWithStyle(fmt.Sprintf(
			"Size: %.0fx%.0f\n" +
			"Bombs: %.0f\n" +
			"Seed: %.0f\n" +
			"First Cell: (%d,%d)\n" +
			"Elapsed Time: %d:%02d:%02d",
			width, height, bombs, seed, firstPos.X, firstPos.Y, et_h, et_m, et_s), fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewButton("Replay?", func() {
			window_main.SetContent(newTitleLayout())
			window_main.Resize(fyne.NewSize(640, 0))
			window_result.Close()
		}),
	)
}

func main() {
	application = app.New()
	window_main = application.NewWindow(TITLE)
	window_main.SetContent(newTitleLayout())
	window_main.Resize(fyne.NewSize(640, 0))
	window_main.SetMaster()
	window_main.ShowAndRun()
}
