package repository

import (
	"context"

	"github.com/m-bromo/my-game-list/internal/domain"
)

type GameRepository interface {
	GetAllGames(ctx context.Context) ([]*domain.Game, error)
}
