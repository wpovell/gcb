package wrapper

import (
	"gcb/config"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Data returned from wrapped block, specifying data to display
type TextData struct {
	text  []string
	color []xgraphics.BGRA
}

func NewTextData() *TextData {
	return &TextData{
		text:  make([]string, 0),
		color: make([]xgraphics.BGRA, 0),
	}
}

// Add text segment of given color to object
func (t *TextData) Color(txt string, color xgraphics.BGRA) *TextData {
	t.text = append(t.text, txt)
	t.color = append(t.color, color)
	return t
}

// Add text segment with default foreground to object
func (t *TextData) Text(txt string) *TextData {
	t.text = append(t.text, txt)
	t.color = append(t.color, config.FG)
	return t
}
