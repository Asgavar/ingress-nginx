//go:build mage

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

var DEBUG bool

func init() {
	DEBUG = false
	debugENV := os.Getenv("MAGE_DEBUG")
	if debugENV == "true" {
		DEBUG = true
	}
}

// CheckArgs should be used to ensure the right command line arguments are
// passed before executing an example.
func CheckArgs(arg ...string) {
	if len(os.Args) < len(arg)+1 {
		ErrorF("Usage: %s %s", os.Args[0], strings.Join(arg, " "))
		os.Exit(1)
	}
}

// CheckIfError should be used to naively panics if an error is not nil.
func CheckIfError(err error, format string, args ...interface{}) {
	if err == nil {
		return
	}

	fmt.Printf("\x1b[31;1m%s ERROR %s %s\x1b[0m\n", timeStamp(), fmt.Sprintf(format, args...), err)
	os.Exit(1)
}

// Info should be used to describe the example commands that are about to run.
func Info(format string, args ...interface{}) {
	fmt.Printf("\x1b[34;1m%s INFO: %s\x1b[0m\n", timeStamp(), fmt.Sprintf(format, args...))
}

func timeStamp() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

// Warning should be used to display a warning
func Warning(format string, args ...interface{}) {
	fmt.Printf("\x1b[36;1m%s WARNING: %s\x1b[0m\n", timeStamp(), fmt.Sprintf(format, args...))
}

// Info should be used to describe the example commands that are about to run.
func Debug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf("\x1b[34;1m%s DEBUG: %s\x1b[0m\n", timeStamp(), fmt.Sprintf(format, args...))
	}
}

// Info should be used to describe the example commands that are about to run.
func ErrorF(format string, args ...interface{}) {
	fmt.Printf("\x1b[31;1m%s ERROR: %s\x1b[0m\n", timeStamp(), fmt.Sprintf(format, args...))
}
