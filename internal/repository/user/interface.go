package repository

import (
	"context"
	"errors"

	"github.com/m-bromo/my-game-list/internal/domain"
)

var (
	ErrUserNotFound = errors.New("user was not found")
)

type UserRepository interface {
	Save(ctx context.Context, user *domain.User) error
	GetByEmail(ctx context.Context, email string) (*domain.User, error)
}
