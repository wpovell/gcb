package config

import (
	"gcb/color"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

const (
	FontFn   = "/home/wpovell/.local/share/fonts/FiraCode-Regular.ttf"
	FontSize = 12
	FontDpi  = 72
	BarH     = 27
)

var FG, BG xgraphics.BGRA

func init() {
	BG = color.HexToBGRA("#2E3440")
	FG = color.HexToBGRA("#D8DEE9")
}
