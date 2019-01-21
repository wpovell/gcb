package music

import (
	"fmt"
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"
)

const (
	iconMusic = ""
)

type Music struct {
	w.NoMsg
	spot *Spotify
}

func New(b *bar.Bar) *w.TextW {
	spot := NewSpotify()
	return w.NewTextW(b, &Music{
		spot: spot,
	})
}

func (m *Music) HandleClick(ev bar.ClickEvent) bool {
	switch ev.Button {
	case bar.LeftClick:
		m.spot.Toggle()
	case bar.MiddleClick:
		m.spot.Prev()
	case bar.RightClick:
		m.spot.Next()
	}

	return true
}

func (m *Music) Interval() time.Duration {
	return time.Second
}

func (m *Music) Text() *w.TextData {
	status := m.spot.Status()

	text := iconMusic
	if status != Quit {
		m := m.spot.Metadata()
		text = fmt.Sprintf("%s %s - %s", iconMusic, m.Artist, m.Title)
	}

	color := config.FG
	if status == Playing {
		color = config.Bright
	}

	return w.NewTextData().Color(text, color)
}
