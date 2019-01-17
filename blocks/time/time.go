package time

import (
	"time"

	"gcb/bar"
	"gcb/blocks/wrapper"

	"github.com/BurntSushi/xgbutil/xevent"
)

type Time struct{}

func Create(b *bar.Bar) *wrapper.TextW {
	return wrapper.CreateTextW(b, &Time{})
}

func (t *Time) Handle(ev xevent.ButtonPressEvent) {}

func (t *Time) Interval() time.Duration {
	return time.Now().Truncate(time.Minute).Add(time.Minute).Sub(time.Now())
}

func (t *Time) Text() string {
	return time.Now().Format("Jan 02 03:04 PM")
}
