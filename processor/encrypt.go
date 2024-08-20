package processor

func Encrypt(data string) error {
	passphrase, err := GetPassphrase()
	if err != nil {
		return err
	}
}
