package main

import (
	"alg_chipher/processor"
	"alg_chipher/utils"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
)

var cipherMode = flag.Bool("cipher", false, "Enable chipher mode.")
var decipherMode = flag.Bool("decipher", false, "Enable decipher mode.")

func main() {
	flag.Parse()

	if *cipherMode && *decipherMode {
		utils.HaltOnErr(errors.New("please specify only one mode at  a time"))

	}

	if *decipherMode {
		decipheredBytes, err := processor.Decrypt()
		utils.HaltOnErr(err)
		fmt.Println(string(decipheredBytes))
	} else if *cipherMode {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter full sentence: ")

		sentence, _ := reader.ReadString('\n')
		err := processor.Encrypt(sentence)
		utils.HaltOnErr(err)
	} else {
		utils.HaltOnErr(errors.New("unknow mode"))
	}
}
