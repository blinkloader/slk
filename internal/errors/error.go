package errors

import (
	"fmt"
)

func New(s string) error {
	return fmt.Errorf(s)
}

func Wrap(base, err error) error {
	return fmt.Errorf("%s: %s", base, err)
}
