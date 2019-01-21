package bat

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"gcb/log"
)

var bat_path string

func init() {
	bat_path = fmt.Sprintf("/sys/class/power_supply/%s", BAT_NAME)
}

// Possible battery states
type Status int

const (
	Charging Status = iota
	Discharging
	Unknown
	Full
	Other
)

var states = map[Status]string{
	Unknown:     "Unknown",
	Discharging: "Discharging",
	Charging:    "Charging",
	Full:        "Full",
	Other:       "Other",
}

func (s Status) String() string {
	res := states[s]
	if res != "" {
		return res
	} else {
		return states[Other]
	}
}

// Default battery to report
const BAT_NAME = "BAT0"

// Read int from file
func readInt(fn string) (int, error) {
	raw, err := ioutil.ReadFile(fn)
	if err != nil {
		return 0, err
	}
	dat := strings.TrimSpace(string(raw))
	i, err := strconv.ParseInt(dat, 10, 64)
	return int(i), err
}

// Join `f` to base battery path
func path(f string) string {
	return filepath.Join(bat_path, f)
}

// Get the charge of the battery
func charge() (int, error) {
	// Now
	now, err := readInt(path("energy_now"))
	if err != nil {
		return 0, err
	}

	// Full
	full, err := readInt(path("energy_full"))
	if err != nil {
		return 0, err
	}
	return 100 * now / full, nil
}

// Get status of battery
func status() Status {
	raw, err := ioutil.ReadFile(filepath.Join(bat_path, "status"))
	if err != nil {
		return Unknown
	}
	dat := strings.TrimSpace(string(raw))

	for i, state := range states {
		if strings.EqualFold(dat, state) {
			return Status(i)
		}
	}

	return Other
}

// Combined charge and status information
type BatInfo struct {
	status Status
	charge int
}

func info() BatInfo {
	cur_charge, err := charge()
	log.Fatal(err)

	return BatInfo{
		status: status(),
		charge: cur_charge,
	}
}
