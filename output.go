package log

import (
	"fmt"
	"os"
)

func StdOutOutput(messageString string) {
	fmt.Println(messageString)
}

func StdErrOutput(messageString string) {
	fmt.Fprintln(os.Stderr, messageString)
}
