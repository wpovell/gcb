package bat

import (
	"fmt"
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"
)

type Bat struct {
	w.NoHandle
}

func New(b *bar.Bar) *w.TextW {
	return w.NewTextW(b, new(Bat))
}

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
