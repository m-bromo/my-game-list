package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/m-bromo/my-game-list/internal/domain"
	"github.com/m-bromo/my-game-list/internal/infra/database/sqlc"
)

type sqlcUserRepository struct {
	querier sqlc.Querier
}

func NewSqlcUserRepository(querier sqlc.Querier) UserRepository {
	return &sqlcUserRepository{
		querier: querier,
	}
}

func (r *sqlcUserRepository) Save(ctx context.Context, user *domain.User) error {
	err := r.querier.Save(ctx, sqlc.SaveParams{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Password: user.Password,
	})

	if err != nil {
		return fmt.Errorf("save user: %w", err)
	}

	return nil
}

func (r *sqlcUserRepository) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	sqlcUser, err := r.querier.GetByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("get user by email: %w", ErrUserNotFound)
		}
		return nil, fmt.Errorf("get user by email: %w", ErrUserNotFound)
	}

	return &domain.User{
		ID:       sqlcUser.ID,
		Email:    sqlcUser.Email,
		Username: sqlcUser.Username,
		Password: sqlcUser.Password,
	}, nil
}
