package log

import (
	"github.com/pkg/errors"
)

// Disable logging entirely
const disable = true

// Enable specific tags
var enabled_tags = map[string]bool{
	"stop": false,
	"draw": false,
}

// Log and panic if an error
func Fatal(err error) {
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
}

// Log if any of tags are enabled
func Log(s string, tags ...string) {
	if disable {
		return
	}
	for _, tag := range tags {
		if enabled_tags[tag] {
			println(s)
			return
		}
	}
}
