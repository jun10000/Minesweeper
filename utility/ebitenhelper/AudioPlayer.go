package ebitenhelper

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
)

const SampleRate = 44100

type AudioPlayer struct {
	context *audio.Context
	players map[string]*audio.Player
}

func NewAudioPlayer() *AudioPlayer {
	c := audio.CurrentContext()
	if c == nil {
		c = audio.NewContext(SampleRate)
	}

	return &AudioPlayer {
		context: c,
		players: make(map[string]*audio.Player),
	}
}

func (p *AudioPlayer) Add(name string, data []byte) {
	reader := bytes.NewReader(data)
	stream, _ := mp3.DecodeWithSampleRate(SampleRate, reader)
	player, _ := p.context.NewPlayer(stream)
	p.players[name] = player
}

func (p *AudioPlayer) Play(name string) {
	player, exists := p.players[name]
	if !exists {
		return
	}

	player.Rewind()
	player.Play()
}

func (p *AudioPlayer) Close() {
	for _, player := range p.players {
		player.Close()
	}
}
