package cmd

import (
	"fmt"
	"os"
)

// Simple error check; print error and exit
func ErrCheck(err error, args ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
