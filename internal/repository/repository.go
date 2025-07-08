package repository

type PasswordStorager interface {
}

type Repository struct {
	Passwords PasswordStorager
}

func NewRepository(passwordStorager PasswordStorager) *Repository {
	return &Repository{Passwords: passwordStorager}
}
