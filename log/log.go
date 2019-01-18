package log

import (
	"fmt"

	"github.com/pkg/errors"
)

func Fatal(err error) {
	if err != nil {
		panic(errors.Wrap(err, ""))
	}
}

func Struct(s interface{}) {
	fmt.Printf("%+v\n", s)
}
