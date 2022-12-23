package widget2

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type Slider2 struct {
	widget.Slider

	ValueData binding.Float
}

func NewSlider2(min, max float64) *Slider2 {
	s := &Slider2{}
	s.ExtendBaseWidget(s)
	s.Value = 0
	s.Min = min
	s.Max = max
	s.Step = 1
	s.Orientation = widget.Horizontal
	return s
}

func NewSlider2WithData(min, max float64, data binding.Float) *Slider2 {
	s := NewSlider2(min, max)
	s.Bind(data)
	return s
}

func (s *Slider2) Bind(data binding.Float) {
	s.Slider.Bind(data)
	s.ValueData = data
}

func (s *Slider2) Unbind() {
	s.Slider.Unbind()
	s.ValueData = nil
}

func (s *Slider2) SetMax(max float64) {
	s.Max = max
	if s.Value >= max {
		if s.ValueData != nil {
			s.ValueData.Set(max)
		} else {
			s.Value = max	// Not tested
		}
	}
	s.Refresh()
}
