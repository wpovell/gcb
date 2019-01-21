package wifi

import (
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"
)

const (
	intf = "wlp4s0"
)

type Wifi struct {
	w.NoHandle
}

func New(b *bar.Bar) *w.TextW {
	return w.NewTextW(b, new(Wifi))
}

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
