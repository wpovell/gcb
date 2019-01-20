package music

import (
	"fmt"
	"time"

	"gcb/bar"
	"gcb/blocks/wrapper"

	"github.com/BurntSushi/xgbutil/xevent"
)

type Music struct {
	spot *Spotify
}

func Create(b *bar.Bar) *wrapper.TextW {
	spot := CreateSpotify()
	return wrapper.CreateTextW(b, &Music{
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

func (m *Music) Text() string {
	status := m.spot.Status()
	ret := status.String()
	if status != Quit {
		m := m.spot.Metadata()
		ret = fmt.Sprintf("%s %s - %s", ret, m.Artist, m.Title)
	}
	return ret
}
