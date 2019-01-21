package time

import (
	"time"

	"gcb/bar"
	w "gcb/blocks/wrapper"
)

type Time struct {
	w.NoHandle
}

func New(b *bar.Bar) *w.TextW {
	return w.NewTextW(b, new(Time))
}

func (t *Time) Interval() time.Duration {
	return time.Now().Truncate(time.Minute).Add(time.Minute).Sub(time.Now())
}

func (t *Time) Text() *w.TextData {
	text := time.Now().Format("Jan 02 03:04 PM")
	return w.NewTextData().Text(text)
}
