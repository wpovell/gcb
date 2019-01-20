package wifi

import (
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"

	"github.com/BurntSushi/xgbutil/xevent"
)

const (
	intf = "wlp4s0"
)

type Wifi struct{}

func New(b *bar.Bar) *w.TextW {
	return w.NewTextW(b, &Wifi{})
}

func (wi *Wifi) Handle(ev xevent.ButtonPressEvent) {}

func (wi *Wifi) Interval() time.Duration {
	return time.Second
}

func (wi *Wifi) Text() *w.TextData {
	txt, err := ssid(intf)
	if err != nil {
		return w.NewTextData().Text("No Wifi")
	} else {
		return w.NewTextData().Color(txt, config.Bright)
	}
}
