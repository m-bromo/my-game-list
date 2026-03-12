package repository

import (
	"context"
	"fmt"

	"github.com/m-bromo/my-game-list/internal/domain"
	"github.com/m-bromo/my-game-list/internal/infra/database/sqlc"
)

type sqlcGameRepository struct {
	querier sqlc.Querier
}

func NewSqlcGameRepository(querier sqlc.Querier) GameRepository {
	return &sqlcGameRepository{
		querier: querier,
	}
}

func (r *sqlcGameRepository) GetAllGames(ctx context.Context) ([]*domain.Game, error) {
	sqlcGames, err := r.querier.GetAllGames(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all games: %w", err)
	}

	games := make([]*domain.Game, len(sqlcGames))
	for _, game := range sqlcGames {
		games = append(games, &domain.Game{
			Name:        game.Name,
			Description: game.Description,
			Genre:       game.Description,
			ImageUrl:    game.ImageUrl.String,
		})
	}

	return games, nil
}
