package bar

import (
	"gcb/log"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xevent"
)

// Initialize X server connection
func initX() *xgbutil.XUtil {
	// New connection to X server
	var err error
	X, err := xgbutil.NewConn()
	log.Fatal(err)

	// Main X event loop
	go xevent.Main(X)

	return X
}
