package widget2

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
)

type intEntry struct {
	widget.Entry
}

func NewIntEntry() *intEntry {
	entry := &intEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func NewIntEntryWithData(data binding.String) *intEntry {
	entry := NewIntEntry()
	entry.Bind(data)
	entry.Validator = nil
	return entry
}

func (e *intEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

func (e *intEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseInt(content, 10, 32); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *intEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}
