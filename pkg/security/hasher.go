package security

type Hasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(inputPassword, hashedPassword string) (bool, error)
}
