package processor

import (
	"crypto/cipher"
	"crypto/rand"
)

func Encrypt(data string) error {
	passphrase, err := GetPassphrase()
	if err != nil {
		return err
	}

	salt, err := makeSalt()
	if err != nil {
		return err
	}

	key, err := DeriveKeyFrom(passphrase, salt)
	if err != nil {
		return err
	}

	crypter, err := MakeCrypterFrom(key)
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

func makeNonceFor(crypter cipher.AEAD) ([]byte, error) {
	nonce := make([]byte, crypter.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}
	return nonce, nil
}
