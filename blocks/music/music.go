package music

import (
	"fmt"
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"

	"github.com/BurntSushi/xgbutil/xevent"
)

type Music struct {
	spot *Spotify
}

func Create(b *bar.Bar) *w.TextW {
	spot := CreateSpotify()
	return w.NewTextW(b, &Music{
		spot: spot,
	})
}

func (m *Music) Handle(ev xevent.ButtonPressEvent) {
	switch ev.Detail {
	case bar.LeftClick:
		m.spot.Toggle()
	case bar.MiddleClick:
		m.spot.Prev()
	case bar.RightClick:
		m.spot.Next()
	}
}

func (m *Music) Interval() time.Duration {
	return time.Second
}

func (m *Music) Text() *w.TextData {
	status := m.spot.Status()
	text := status.String()
	if status != Quit {
		m := m.spot.Metadata()
		text = fmt.Sprintf("%s %s - %s", text, m.Artist, m.Title)
	}

	color := config.FG
	if status == Playing {
		color = config.Bright
	}

	return w.NewTextData().Color(text, color)
}
