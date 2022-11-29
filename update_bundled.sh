#!/bin/bash

fyne bundle resource/close_bomb.svg > bundled.go
fyne bundle -append resource/close_none.svg >> bundled.go
fyne bundle -append resource/close_question.svg >> bundled.go
fyne bundle -append resource/open_bomb.svg >> bundled.go
fyne bundle -append resource/open_near1.svg >> bundled.go
fyne bundle -append resource/open_near2.svg >> bundled.go
fyne bundle -append resource/open_near3.svg >> bundled.go
fyne bundle -append resource/open_near4.svg >> bundled.go
fyne bundle -append resource/open_near5.svg >> bundled.go
fyne bundle -append resource/open_near6.svg >> bundled.go
fyne bundle -append resource/open_near7.svg >> bundled.go
fyne bundle -append resource/open_near8.svg >> bundled.go
fyne bundle -append resource/open_none.svg >> bundled.go
