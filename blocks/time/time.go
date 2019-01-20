package time

import (
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"

	"github.com/BurntSushi/xgbutil/xevent"
)

type Time struct{}

func Create(b *bar.Bar) *w.TextW {
	return w.NewTextW(b, &Time{})
}

func (t *Time) Handle(ev xevent.ButtonPressEvent) {}

func (t *Time) Interval() time.Duration {
	return time.Now().Truncate(time.Minute).Add(time.Minute).Sub(time.Now())
}

func (t *Time) Text() *w.TextData {
	text := time.Now().Format("Jan 02 03:04 PM")
	return w.NewTextData().Text(text)
}
