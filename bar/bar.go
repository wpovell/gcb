package bar

import (
	"image"

	"gcb/config"
	"gcb/log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

type Align int

const (
	Left   Align = 0
	Center Align = 1
	Right  Align = 2
)

type Bar struct {
	// Graphics
	X   *xgbutil.XUtil
	win *xwindow.Window
	img *xgraphics.Image

	// Block lists
	blocks [][]Block

	bg, fg xgraphics.BGRA

	w, h   int
	Redraw chan Block
}

func Create() *Bar {
	bar := new(Bar)
	var err error

	X := initX()
	bar.X = X

	bar.win, err = xwindow.Generate(X)
	log.Fatal(err)

	scr := X.Screen()
	scrW, scrH := int(scr.WidthInPixels), int(scr.HeightInPixels)
	x, y := 0, scrH-config.BarH
	w, h := scrW, config.BarH

	// Bar window
	bar.win.Create(
		X.RootWin(),
		x, y, w, h,
		xproto.CwBackPixel|xproto.CwEventMask,
		0x000000,
		xproto.EventMaskButtonPress,
	)
	log.Fatal(err)

	// EWMH
	err = ewmh.WmWindowTypeSet(X, bar.win.Id,
		[]string{"_NET_WM_WINDOW_TYPE_DOCK"})
	log.Fatal(err)

	err = ewmh.WmStateSet(X, bar.win.Id,
		[]string{"_NET_WM_STATE_STICKY"})
	log.Fatal(err)

	err = ewmh.WmDesktopSet(X, bar.win.Id, ^uint(0))
	log.Fatal(err)

	err = ewmh.WmNameSet(X, bar.win.Id, "bar")
	log.Fatal(err)

	// Map
	bar.win.Map()
	bar.win.Move(x, y)

	// Image
	bar.img = xgraphics.New(X, image.Rect(0, 0, w, h))
	err = bar.img.XSurfaceSet(bar.win.Id)
	log.Fatal(err)
	bar.img.XDraw()

	bar.w = w
	bar.h = h

	bar.Redraw = make(chan Block)
	bar.blocks = make([][]Block, 3)
	for i := 0; i < 3; i++ {
		bar.blocks[i] = make([]Block, 0)
	}

	// Event handler
	xevent.ButtonPressFun(func(_ *xgbutil.XUtil, ev xevent.ButtonPressEvent) {
		bar.findBlock(int(ev.EventX)).Handle(ev)
	}).Connect(X, bar.win.Id)

	return bar
}
