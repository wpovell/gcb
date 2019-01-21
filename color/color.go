package color

import (
	"encoding/hex"
	"strings"

	"gcb/log"

	"github.com/BurntSushi/xgbutil/xgraphics"
)

// Convert 6 character hex string to BGRA
func HexToBGRA(h string) xgraphics.BGRA {
	h = strings.Replace(h, "#", "", 1)
	d, err := hex.DecodeString(h)
	log.Fatal(err)

	return xgraphics.BGRA{B: d[2], G: d[1], R: d[0], A: 0xFF}
}
