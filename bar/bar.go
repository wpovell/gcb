package bar

import (
	"image"
	"sync"

	"gcb/config"
	"gcb/log"

	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
)

// Block alignment on bar
type Align int

const (
	Left   Align = 0
	Center Align = 1
	Right  Align = 2
)

type Bar struct {
	sync.Mutex

	// Graphics
	X   *xgbutil.XUtil
	win *xwindow.Window
	img *xgraphics.Image

	// Block lists
	blocks map[Block]*BlockState
	align  [][]Block

	bg, fg xgraphics.BGRA

	w, h   int
	x, y   int
	Redraw chan DrawState
}

// Current data and placement of a displayed block
type BlockState struct {
	state DrawState
	start int
}

func (s *BlockState) Contains(x int) bool {
	return s.start < x && x < s.start+s.state.Width()
}

func (s *BlockState) Draw(x int, img *xgraphics.Image) {
	s.start = x
	s.state.Draw(x, img)
}

// Initialize X stuff
func (b *Bar) xinit() {
	// New connection to X server
	var err error
	X, err := xgbutil.NewConn()
	log.Fatal(err)
	b.X = X

	// Main X event loop
	go xevent.Main(X)

	// Create X window
	b.win, err = xwindow.Generate(X)
	log.Fatal(err)

	// Calculate dimensions
	scr := X.Screen()
	scrW, scrH := 1920, int(scr.HeightInPixels)
	b.x, b.y = 0, scrH-config.BarH
	b.w, b.h = scrW, config.BarH

	// Bar window
	b.win.Create(
		X.RootWin(),
		b.x, b.y, b.w, b.h,
		xproto.CwBackPixel|xproto.CwEventMask,
		0x000000,
		xproto.EventMaskButtonPress,
	)
	log.Fatal(err)

	// EWMH
	err = ewmh.WmWindowTypeSet(X, b.win.Id,
		[]string{"_NET_WM_WINDOW_TYPE_DOCK"})
	log.Fatal(err)

	err = ewmh.WmStateSet(X, b.win.Id,
		[]string{"_NET_WM_STATE_STICKY"})
	log.Fatal(err)

	err = ewmh.WmDesktopSet(X, b.win.Id, ^uint(0))
	log.Fatal(err)

	err = ewmh.WmNameSet(X, b.win.Id, "bar")
	log.Fatal(err)

	// Map
	b.win.Map()
	b.win.Move(b.x, b.y)

	// Image
	b.img = xgraphics.New(X, image.Rect(0, 0, b.w, b.h))
	err = b.img.XSurfaceSet(b.win.Id)
	log.Fatal(err)
	b.img.XDraw()
}

// Create new bar instance
func Create() *Bar {
	bar := new(Bar)

	bar.xinit()

	bar.Redraw = make(chan DrawState)
	bar.blocks = make(map[Block]*BlockState)
	bar.align = make([][]Block, 3)
	for i := 0; i < 3; i++ {
		bar.align[i] = make([]Block, 0)
	}

	// Event handler
	xevent.ButtonPressFun(func(_ *xgbutil.XUtil, ev xevent.ButtonPressEvent) {
		block := bar.findBlock(int(ev.EventX))
		if block != nil {
			block.Handle(ev)
		}
	}).Connect(bar.X, bar.win.Id)

	return bar
}
