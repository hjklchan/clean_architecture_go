package value_object

type PasswordHasher interface {
	Hash(string) (string, error)
	Compare(plainText, hashedText string) (bool, error)
}

type Password struct {
	hash string
}

func NewPasswordFromHash(hash string) *Password {
	return &Password{
		hash: hash,
	}
}

func NewPasswordFromPlainText(value string, hasher PasswordHasher) (*Password, error) {
	hashed, err := hasher.Hash(value)
	if err != nil {
		return nil, err
	}

	return &Password{
		hash: hashed,
	}, nil
}

func (p *Password) IsMatch(plainText string, hasher PasswordHasher) (bool, error) {
	return hasher.Compare(plainText, p.hash)
}

func (p *Password) Value() string {
	return p.hash
}
