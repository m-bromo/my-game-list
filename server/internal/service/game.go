package service

import (
	"context"
	"fmt"

	"github.com/m-bromo/my-game-list/internal/repository"
	"github.com/m-bromo/my-game-list/internal/web/models"
)

type GameService interface {
	GetAll(ctx context.Context) ([]*models.GameOutput, error)
}

type gameService struct {
	gameRepository repository.GameRepository
}

func NewGameService(gameRepository repository.GameRepository) GameService {
	return &gameService{
		gameRepository: gameRepository,
	}
}

func (s *gameService) GetAll(ctx context.Context) ([]*models.GameOutput, error) {
	games, err := s.gameRepository.GetAllGames(ctx)
	if err != nil {
		return nil, fmt.Errorf("get all: %w", err)
	}

	gamesOutput := make([]*models.GameOutput, len(games))
	for _, game := range gamesOutput {
		gamesOutput = append(gamesOutput, &models.GameOutput{
			Name:        game.Name,
			Description: game.Description,
			Genre:       game.Genre,
			ImageUrl:    game.ImageUrl,
		})
	}

	return gamesOutput, nil
}
