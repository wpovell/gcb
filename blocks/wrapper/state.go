package wrapper

import (
	"image"

	"gcb/bar"
	"gcb/text"

	"github.com/BurntSushi/xgbutil/xgraphics"
	"golang.org/x/image/font"
)

// State passed to bar to request a redraw
type TextWState struct {
	txt    *TextData
	width  int
	block  *TextW
	drawer *font.Drawer
}

// Source of the request
func (s *TextWState) Source() bar.Block {
	return s.block
}

// Width of the text data
func (s *TextWState) Width() int {
	return s.width
}

// Draw state on the passed `img` at `x`
func (s *TextWState) Draw(x int, img *xgraphics.Image) {
	s.drawer.Dst = img
	s.drawer.Dot = text.Point(x)

	for i, text := range s.txt.text {
		color := s.txt.color[i]
		s.drawer.Src = image.NewUniform(color)
		s.drawer.DrawString(text)
	}
}
