package time

import (
	"time"

	"gcb/bar"
	"gcb/text"

	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"golang.org/x/image/font"
)

type Time struct {
	bar    *bar.Bar
	drawer *font.Drawer
	txt    string
}

func Create(b *bar.Bar) *Time {
	return &Time{
		bar:    b,
		drawer: text.Drawer(),
	}
}

func (t *Time) Handle(ev xevent.ButtonPressEvent) {}
func (t *Time) Draw(x int, img *xgraphics.Image) {
	t.drawer.Dst = img
	t.drawer.Dot = text.Point(x)
	t.drawer.DrawString(t.txt)
}

func (t *Time) Width() int {
	return font.MeasureString(t.drawer.Face, t.txt).Ceil()
}

func timeUntilMin() time.Duration {
	return time.Now().Truncate(time.Minute).Add(time.Minute).Sub(time.Now())
}

func (t *Time) Start() {
	go func() {
		for {
			t.txt = time.Now().Format("Jan 02 03:04 PM")
			t.bar.Redraw <- t
			wake := time.Now().Truncate(time.Minute).Add(time.Minute)
			time.Sleep(time.Until(wake))
		}
	}()
}
