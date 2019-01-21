package log

import (
	"fmt"
	"runtime"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type Logger struct {
	disabled     bool
	show_src     bool
	enabled_tags map[string]bool
}

var global_logger Logger

// Mutex used to garuntee exclusive access to stdout
var writeMtx sync.Mutex

// Initialize global logger
func init() {
	global_logger = Logger{
		disabled: false,
		show_src: true,
		enabled_tags: map[string]bool{
			"stop": true,
			"draw": false,
			"warn": true,
			"ipc":  false,
		},
	}
}

// Check if any of `tags` are enabled
func (l Logger) is_enabled(tags []string) bool {
	if l.disabled {
		return false
	}
	for _, tag := range tags {
		if l.enabled_tags[tag] {
			return true
		}
	}
	return false
}

// Log and panic if an error
func Fatal(err error) {
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
}

// Write out logline for `msg`
func (l Logger) prompt(msg string) {
	writeMtx.Lock()
	defer writeMtx.Unlock()

	if !l.show_src {
		println(msg)
		return
	}
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		println(msg)
		return
	} else {
		parts := strings.Split(file, "/")
		for i, part := range parts {
			if part == "gcb" {
				parts = parts[i+1:]
				break
			}
		}
		file = strings.Join(parts, "/")
		fmt.Printf("[%s:%d] %s\n", file, line, msg)
	}
}

// Log `msg` if any of `tags` are enabled
func (l Logger) Log(msg string, tags ...string) {
	if l.is_enabled(tags) {
		l.prompt(msg)
	}
}

// Log using global logger
func Log(msg string, tags ...string) {
	global_logger.Log(msg, tags...)
}
