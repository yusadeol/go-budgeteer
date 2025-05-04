package mock

type PasswordHasher struct{}

func NewPasswordHasher() *PasswordHasher {
	return &PasswordHasher{}
}

func (m PasswordHasher) Hash(password string) (string, error) {
	return "", nil
}

func (m PasswordHasher) Verify(password, hash string) (bool, error) {
	return true, nil
}
