package log

import (
	"fmt"

	"github.com/pkg/errors"
)

const disable = true

var enabled_tags = map[string]bool{
	"stop": false,
}

func Fatal(err error) {
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
}

func Struct(s interface{}) {
	fmt.Printf("%+v\n", s)
}

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
