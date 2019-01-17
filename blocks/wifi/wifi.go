package wifi

import (
	"time"

	"gcb/bar"
	"gcb/blocks/wrapper"

	"github.com/BurntSushi/xgbutil/xevent"
)

const (
	intf = "wlp4s0"
)

type Wifi struct{}

func Create(b *bar.Bar) *wrapper.TextW {
	return wrapper.CreateTextW(b, &Wifi{})
}

func (w *Wifi) Handle(ev xevent.ButtonPressEvent) {}

func (w *Wifi) Interval() time.Duration {
	return time.Second
}

func (w *Wifi) Text() string {
	txt, err := ssid(intf)
	if err != nil {
		return "No Wifi"
	} else {
		return txt
	}
}
