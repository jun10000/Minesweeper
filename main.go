package main

import (
	"jun10000.github.io/minesweeper/container2"
	"jun10000.github.io/minesweeper/resource"
	"jun10000.github.io/minesweeper/utility"
	"jun10000.github.io/minesweeper/utility/ebitenhelper"
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
	audioplayer *ebitenhelper.AudioPlayer

	width = float64(10)
	height = float64(10)
	bombs = float64(10)
	seed = float64(0)
)

func newSeed() float64 {
	return float64(time.Now().Nanosecond())
}

func playAudio(r *fyne.StaticResource) {
	audioplayer.Play(r.StaticName)
}

func newTitleLayout() *fyne.Container {
	seed = newSeed()
	
	width_data := binding.BindFloat(&width)
	height_data := binding.BindFloat(&height)
	bombs_data := binding.BindFloat(&bombs)
	seed_data := binding.BindFloat(&seed)

	bombs_slider := widget2.NewSlider2WithData(2, 9990, bombs_data)
	func_updateMaxBombs := func() {
		max := width * height - 10
		bombs_slider.SetMax(max)
	}
	width_data.AddListener(binding.NewDataListener(func_updateMaxBombs))
	height_data.AddListener(binding.NewDataListener(func_updateMaxBombs))

	return container.NewVBox(
		widget.NewLabelWithStyle(TITLE, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		container2.NewOneColumnMax(1,
			widget.NewLabel("Width"),
			widget2.NewSlider2WithData(4, 50, width_data),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(width_data, "%.0f")),
		),
		container2.NewOneColumnMax(1,
			widget.NewLabel("Height"),
			widget2.NewSlider2WithData(4, 50, height_data),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(height_data, "%.0f")),
		),
		container2.NewOneColumnMax(1,
			widget.NewLabel("Bombs"),
			bombs_slider,
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(bombs_data, "%.0f")),
		),
		container.NewHBox(
			widget.NewLabel("Seed"),
			widget2.NewIntEntryWithData(binding.FloatToStringWithFormat(seed_data, "%.0f")),
			widget.NewButton("Refresh", func() {
				seed_data.Set(newSeed())
			}),
		),
		widget.NewButton("START", func() {
			playAudio(resource.StartMp3)
			window_main.SetContent(newGameLayout())
			window_main.Resize(fyne.NewSize(0, 0))
		}),
	)
}

func newGameLayout() *fyne.Container {
	return widget2.NewMSTable(int(width), int(height), int(bombs), int64(seed),
		func(elapsedTime time.Duration, firstPos utility.Position) {
			playAudio(resource.ClearMp3)
			window_result = application.NewWindow("Result")
			window_result.SetContent(newClearLayout(elapsedTime, firstPos))
			window_result.Resize(fyne.NewSize(300, 0))
			window_result.Show()
		},
		func(elapsedTime time.Duration, firstPos utility.Position) {
			playAudio(resource.GameoverMp3)
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
			"Field: %.0fx%.0f (%.0f Bombs)\n" +
			"Seed: %.0f\n" +
			"First Cell: (%d,%d)\n" +
			"Elapsed Time: %d:%02d:%02d",
			width, height, bombs, seed, firstPos.X + 1, firstPos.Y + 1, et_h, et_m, et_s), fyne.TextAlignCenter, fyne.TextStyle{}),
		widget.NewButton("Return to Top", func() {
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
			"Field: %.0fx%.0f (%.0f Bombs)\n" +
			"Seed: %.0f\n" +
			"First Cell: (%d,%d)\n" +
			"Elapsed Time: %d:%02d:%02d",
			width, height, bombs, seed, firstPos.X + 1, firstPos.Y + 1, et_h, et_m, et_s), fyne.TextAlignCenter, fyne.TextStyle{}),
		container.NewGridWithColumns(2,
			widget.NewButton("Retry", func() {
				playAudio(resource.Start2Mp3)
				window_main.SetContent(newGameLayout())
				window_main.Resize(fyne.NewSize(0, 0))
				window_result.Close()
			}),
			widget.NewButton("Return to Top", func() {
				window_main.SetContent(newTitleLayout())
				window_main.Resize(fyne.NewSize(640, 0))
				window_result.Close()
			}),
		),
	)
}

func main() {
	audioList := []*fyne.StaticResource {
		resource.StartMp3,
		resource.Start2Mp3,
		resource.ClearMp3,
		resource.GameoverMp3,
	}

	audioplayer = ebitenhelper.NewAudioPlayer()
	for _, r := range audioList {
		audioplayer.Add(r.StaticName, r.StaticContent)
	}

	application = app.New()
	window_main = application.NewWindow(TITLE)
	window_main.SetContent(newTitleLayout())
	window_main.Resize(fyne.NewSize(640, 0))
	window_main.SetMaster()
	window_main.ShowAndRun()

	audioplayer.Close()
}
