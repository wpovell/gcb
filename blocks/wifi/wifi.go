package wifi

import (
	"fmt"
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
	"gcb/config"
)

const (
	intf = "wlp4s0"
)

const (
	iconWifi   = "直"
	iconNoWifi = "睊"
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
		return w.NewTextData().Text(iconNoWifi)
	} else {
		return w.NewTextData().Color(fmt.Sprintf("%s %s", iconWifi, txt), config.Bright)
	}
}
