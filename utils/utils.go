package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var errorLoger = log.New(os.Stderr, "ERROR: ", log.LstdFlags|log.Lshortfile)

func HaltOnErr(err error, messages ...string) {
	if err == nil {
		return
	}

	message := "An error occured"
	if len(messages) > 0 {
		message = fmt.Sprintf("%s: %s", message, strings.Join(messages, " "))
	}

	errorLoger.Printf("%s: %v", message, err)
	os.Exit(1)
}
