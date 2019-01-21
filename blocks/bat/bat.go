package bat

import (
	"fmt"
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"
)

var iconBat = []string{
	"",
	"",
	"",
	"",
	"",
}

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
	} else if bat.charge <= 10 {
		color = config.Red
	}

	ind := int(float64(bat.charge) / 100.0 * float64(len(iconBat)-1))
	text := fmt.Sprintf("%s %d%%", iconBat[ind], bat.charge)
	return w.NewTextData().Color(text, color)
}
