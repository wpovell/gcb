package text

import (
	"image"
	"io/ioutil"

	"gcb/config"
	"gcb/log"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"
)

// Global font to use
var face *truetype.Font

// Load font to be used by bar
func init() {
	var err error

	fontBytes, err := ioutil.ReadFile(config.FontFn)
	log.Fatal(err)

	face, err = freetype.ParseFont(fontBytes)
	log.Fatal(err)
}

// Get position for text
func Point(x int) fixed.Point26_6 {
	// This is a maybe-correct formula to center the text vertically
	return fixed.P(x, (config.FontSize/2)+(config.BarH/2)-2)
}

// Return a new drawer for the global bar font
func Drawer() *font.Drawer {
	return &font.Drawer{
		Src: image.NewUniform(config.FG),
		Face: truetype.NewFace(face, &truetype.Options{
			Size:    config.FontSize,
			DPI:     config.FontDpi,
			Hinting: font.HintingFull,
		}),
		Dot: Point(0),
	}
}
