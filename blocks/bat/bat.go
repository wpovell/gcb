package bat

import (
	"fmt"
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"

	"github.com/BurntSushi/xgbutil/xevent"
)

type Bat struct{}

func Create(b *bar.Bar) *w.TextW {
	return w.NewTextW(b, &Bat{})
}

func (b *Bat) Handle(ev xevent.ButtonPressEvent) {}

func (b *Bat) Interval() time.Duration {
	return time.Second
}

func (b *Bat) Text() *w.TextData {
	bat := info()

	color := config.FG
	if bat.state == Full || bat.state == Charging {
		color = config.Bright
	}

	text := fmt.Sprintf("%d%%", bat.charge)
	return w.NewTextData().Color(text, color)
}
