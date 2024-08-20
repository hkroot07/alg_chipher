package main

import (
	"errors"
	"flag"
)

var cipherMode = flag.Bool("cipher", false, "Enable chipher mode.")
var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")

func main() {
	flag.Parse()

	if *cipherMode && *decipherMode {
		utils.HaltOnErr(errors.New("please specify only one mode at  a time"))

	}

	if *decipherMode {

	} else {

	}
}
