package config

import (
	"gcb/color"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Font config
const (
	FontFn   = "./resource/fura_code_nerd.ttf"
	FontSize = 12
	FontDpi  = 72
)

// Dimensions
const (
	BarH    = 27
	Padding = 10
)

// Colors used by bar
var FG, BG, Bright, Red xgraphics.BGRA

func init() {
	BG = color.HexToBGRA("#2E3440")
	FG = color.HexToBGRA("#D8DEE9")
	Bright = color.HexToBGRA("#EBCB8B")
	Red = color.HexToBGRA("#bb0000")
}
