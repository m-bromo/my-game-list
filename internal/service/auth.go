package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/m-bromo/my-game-list/internal/domain"
	repository "github.com/m-bromo/my-game-list/internal/repository/user"
	"github.com/m-bromo/my-game-list/internal/web/models"
	"github.com/m-bromo/my-game-list/pkg/security"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
)

type AuthService interface {
	Register(ctx context.Context, input *models.RegisterInput) error
}

type authService struct {
	userRepository repository.UserRepository
	hasher         security.Hasher
}

func NewService(userRepository repository.UserRepository, hasher security.Hasher) AuthService {
	return &authService{
		userRepository: userRepository,
		hasher:         hasher,
	}
}

func (s *authService) Register(ctx context.Context, input *models.RegisterInput) error {
	err := s.userRepository.GetByEmail(ctx, input.Email)
	if err != nil && errors.Is(err, repository.ErrUserNotFound) {
		return fmt.Errorf("register user: %w", err)
	}

	hashedPassword, err := s.hasher.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("register user: %w", err)
	}

	user := &domain.User{
		ID:       uuid.New(),
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := s.userRepository.Save(ctx, user); err != nil {
		return fmt.Errorf("register user: %w", err)
	}

	return nil
}
