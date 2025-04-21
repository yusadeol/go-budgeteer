package adapter

import "golang.org/x/crypto/bcrypt"

type BcryptHasher struct {
	Cost int
}

func NewBcryptHasher(cost int) *BcryptHasher {
	return &BcryptHasher{cost}
}

func (p BcryptHasher) Hash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), p.Cost)
	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func (p BcryptHasher) Verify(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false, err
	}

	return true, nil
}
