package log

import (
	"github.com/pkg/errors"
)

func Fatal(err error) {
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
}
