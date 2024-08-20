package processor

import "crypto/rand"

func Encrypt(data string) error {
	passphrase, err := GetPassphrase()
	if err != nil {
		return err
	}

	salt, err := makeSalt()
	if err != nil {
		return err
	}
}

// generate salt
func makeSalt() ([]byte, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return nil, err
	}
	return salt, nil
}
