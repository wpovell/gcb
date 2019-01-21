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
	// Can probably be connected to a wifi w/o a name but not sure
	// how else to differentiate
	if err != nil || txt == "" {
		return w.NewTextData().Text(iconNoWifi)
	} else {
		return w.NewTextData().Color(fmt.Sprintf("%s %s", iconWifi, txt), config.Bright)
	}
}
