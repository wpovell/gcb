package music

import (
	"errors"
	"fmt"

	"gcb/log"

	"github.com/godbus/dbus"
)

type Spotify struct {
	sdbus dbus.BusObject
}

func CreateSpotify() *Spotify {
	conn, err := dbus.SessionBus()
	log.Fatal(err)
	obj := conn.Object("org.mpris.MediaPlayer2.spotify", "/org/mpris/MediaPlayer2")
	return &Spotify{
		sdbus: obj,
	}
}

// Metadata
func (s *Spotify) Metadata() Metadata {
	song, err := s.sdbus.GetProperty("org.mpris.MediaPlayer2.Player.Metadata")
	log.Fatal(err)

	songData := song.Value().(map[string]dbus.Variant)
	m := Metadata{}
	// Only ever one artist
	m.Artist = songData["xesam:artist"].Value().([]string)[0]
	m.Title = songData["xesam:title"].Value().(string)

	return m
}

type Metadata struct {
	Artist string
	Title  string
}

// Status
func (s *Spotify) Status() Status {
	raw_status, err := s.sdbus.GetProperty("org.mpris.MediaPlayer2.Player.PlaybackStatus")
	if err != nil {
		return Quit
	}

	status := raw_status.Value()
	for i, name := range statuses {
		if name == status {
			return Status(i)
		}
	}
	log.Fatal(errors.New(fmt.Sprintf("Unknown status: '%s'", raw_status)))
	return 0
}

type Status int

const (
	Quit Status = iota
	Playing
	Paused
)

var statuses = map[Status]string{
	Quit:    "Quit",
	Playing: "Playing",
	Paused:  "Paused",
}

func (m Status) String() string {
	return statuses[m]
}
