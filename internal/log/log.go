package log

import (
	"fmt"
	"os"
)

func Fatal(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	fmt.Printf(format, v...)
	os.Exit(1)
}
