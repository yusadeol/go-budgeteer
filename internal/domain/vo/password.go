package vo

import "errors"

var (
	ErrPasswordIsNotEqual = errors.New("password not equal")
)

type PasswordHasher interface {
	Hash(password string) (string, error)
	Verify(password, hash string) (bool, error)
}

type Password struct {
	hasher PasswordHasher
	Value  string
}

func NewPassword(hasher PasswordHasher) *Password {
	return &Password{hasher: hasher}
}

func NewPasswordFromHash(hasher PasswordHasher, password string) *Password {
	return &Password{
		hasher: hasher,
		Value:  password,
	}
}

func (p *Password) Parse(password string) error {
	hash, err := p.hasher.Hash(password)
	if err != nil {
		return err
	}

	p.Value = hash

	return nil
}

func (p *Password) Compare(password string) error {
	isEqual, err := p.hasher.Verify(password, p.Value)
	if err != nil {
		return err
	}

	if !isEqual {
		return ErrPasswordIsNotEqual
	}

	return nil
}
