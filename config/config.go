package config

import (
	"gcb/color"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

const (
	FontFn    = "/home/wpovell/.local/share/fonts/FiraCode-Regular.ttf"
	FontSize  = 12
	FontDpi   = 72
	BarH = 27
)

var FG, BG xgraphics.BGRA

func init() {
	FG = color.HexToBGRA("#2E3440")
	BG = color.HexToBGRA("#D8DEE9")
}
