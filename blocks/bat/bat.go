package bat

import (
	"fmt"
	"time"

	"gcb/bar"
	"gcb/blocks/wrapper"

	"github.com/BurntSushi/xgbutil/xevent"
)

type Bat struct{}

func Create(b *bar.Bar) *wrapper.TextW {
	return wrapper.CreateTextW(b, &Bat{})
}

func (b *Bat) Handle(ev xevent.ButtonPressEvent) {}

func (b *Bat) Interval() time.Duration {
	return time.Second
}

func (b *Bat) Text() string {
	bat := info()
	if bat.state == Charging {
		return fmt.Sprintf("Charging %d%%", bat.charge)
	} else {
		return fmt.Sprintf("Discharging %d%%", bat.charge)
	}
}
