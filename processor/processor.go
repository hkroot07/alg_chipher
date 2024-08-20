package processor

import (
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/term"
)

func DeriveKeyFrom(passphrase, salt []byte) ([]byte, error) {
	key := argon2.IDKey(passphrase, salt, 1, 64*1024, 2, 32)
	return key, nil
}

func GetPassphrase() ([]byte, error) {
	println("Enter passphrase:")
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))

	if err != nil {
		return nil, fmt.Errorf("failed to grab passphrase: %w", err)
	}

	passphrase, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return nil, fmt.Errorf("failed to grab passphrase: %w", err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	return passphrase, nil
}
