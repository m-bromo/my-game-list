package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/m-bromo/my-game-list/internal/domain"
	"github.com/m-bromo/my-game-list/internal/repository"
	"github.com/m-bromo/my-game-list/internal/web/models"
	"github.com/m-bromo/my-game-list/pkg/security"
)

var (
	ErrEmailAlreadyRegistered = errors.New("email already registered")
	ErrEmailNotRegistered     = errors.New("the email was not registered")
	ErrInvalidCredentials     = errors.New("the credentials provided were not valid")
)

type AuthService interface {
	Register(ctx context.Context, input *models.RegisterInput) error
	Login(ctx context.Context, input *models.LoginInput) error
}

type authService struct {
	userRepository repository.UserRepository
	hasher         security.Hasher
}

func NewUserService(userRepository repository.UserRepository, hasher security.Hasher) AuthService {
	return &authService{
		userRepository: userRepository,
		hasher:         hasher,
	}
}

func (s *authService) Register(ctx context.Context, input *models.RegisterInput) error {
	_, err := s.userRepository.GetByEmail(ctx, input.Email)
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

func (s *authService) Login(ctx context.Context, input *models.LoginInput) error {
	foundUser, err := s.userRepository.GetByEmail(ctx, input.Email)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return fmt.Errorf("login: %w", ErrEmailNotRegistered)
		}

		return fmt.Errorf("login: %w", err)
	}

	ok, err := s.hasher.VerifyPassword(input.Password, foundUser.Password)
	if err != nil {
		return fmt.Errorf("login: %w", err)
	}

	if !ok {
		return fmt.Errorf("login: %w", ErrInvalidCredentials)
	}

	return nil
}
