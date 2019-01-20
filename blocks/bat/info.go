package bat

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"

	"gcb/log"
)

type State int

const (
	Charging State = iota
	Discharging
	Unknown
	Full
	Other
)

var states = map[State]string{
	Unknown:     "Unknown",
	Discharging: "Discharging",
	Charging:    "Charging",
	Full:        "Full",
	Other:       "Other",
}

func (s State) String() string {
	res := states[s]
	if res != "" {
		return res
	} else {
		return states[Other]
	}
}

const BAT_NAME = "BAT0"

type BatInfo struct {
	state  State
	charge int
}

var bat_path string

func init() {
	bat_path = fmt.Sprintf("/sys/class/power_supply/%s", BAT_NAME)
}

func info() BatInfo {
	cur_charge, err := charge()
	log.Fatal(err)

	return BatInfo{
		state:  state(),
		charge: cur_charge,
	}
}

func readInt(fn string) (int, error) {
	raw, err := ioutil.ReadFile(fn)
	if err != nil {
		return 0, err
	}
	dat := strings.TrimSpace(string(raw))
	i, err := strconv.ParseInt(dat, 10, 64)
	return int(i), err
}

func path(f string) string {
	return filepath.Join(bat_path, f)
}

func charge() (int, error) {
	now, err := readInt(path("energy_now"))
	if err != nil {
		return 0, err
	}
	full, err := readInt(path("energy_full"))
	if err != nil {
		return 0, err
	}
	return 100 * now / full, nil
}

func state() State {
	raw, err := ioutil.ReadFile(filepath.Join(bat_path, "status"))
	if err != nil {
		return Unknown
	}
	dat := strings.TrimSpace(string(raw))

	for i, state := range states {
		if strings.EqualFold(dat, state) {
			return State(i)
		}
	}

	return Other
}
