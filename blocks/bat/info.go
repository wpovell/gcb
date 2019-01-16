package bat

type State = int

const (
	Charging State = iota
	Draining State = iota
)

type BatInfo struct {
	state  State
	charge int
}

func info() BatInfo {
	return BatInfo{
		state:  Charging,
		charge: 100,
	}
}
