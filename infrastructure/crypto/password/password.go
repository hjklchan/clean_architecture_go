package password

import "golang.org/x/crypto/bcrypt"

type CryptoPasswordHasher struct{}

func (CryptoPasswordHasher) Hash(plainText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)

	return string(hash), err
}

func (CryptoPasswordHasher) Compare(plainText, hashedText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedText), []byte(plainText))
	if err != nil {
		return false, err
	}

	return true, nil
}
