#!/bin/bash

cd `dirname $0`

fyne bundle --package "resource" --prefix "" close_bomb.png > bundled.go
fyne bundle --append --prefix "" close_none.png >> bundled.go
fyne bundle --append --prefix "" close_question.png >> bundled.go
fyne bundle --append --prefix "" open_bomb.png >> bundled.go
fyne bundle --append --prefix "" open_near1.png >> bundled.go
fyne bundle --append --prefix "" open_near2.png >> bundled.go
fyne bundle --append --prefix "" open_near3.png >> bundled.go
fyne bundle --append --prefix "" open_near4.png >> bundled.go
fyne bundle --append --prefix "" open_near5.png >> bundled.go
fyne bundle --append --prefix "" open_near6.png >> bundled.go
fyne bundle --append --prefix "" open_near7.png >> bundled.go
fyne bundle --append --prefix "" open_near8.png >> bundled.go
fyne bundle --append --prefix "" open_none.png >> bundled.go

fyne bundle --append --prefix "" start.mp3 >> bundled.go
fyne bundle --append --prefix "" clear.mp3 >> bundled.go
fyne bundle --append --prefix "" gameover.mp3 >> bundled.go
