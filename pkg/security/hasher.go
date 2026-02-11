package security

type Hasher interface {
	HashPassword(password string) (string, error)
}
