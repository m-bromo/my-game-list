package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/my-game-list/config"
	"github.com/m-bromo/my-game-list/internal/infra/database"
	"github.com/m-bromo/my-game-list/internal/infra/database/sqlc"
	"github.com/m-bromo/my-game-list/internal/repository"
	"github.com/m-bromo/my-game-list/internal/service"
	"github.com/m-bromo/my-game-list/internal/web/handler"
	"github.com/m-bromo/my-game-list/internal/web/routes"
	"github.com/m-bromo/my-game-list/pkg/logging"
	"github.com/m-bromo/my-game-list/pkg/security"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	logger := logging.NewLogger(config)
	db, err := database.NewPostgresConnection(config)
	if err != nil {
		log.Fatal(err)
	}

	g := gin.New()

	querier := sqlc.New(db)

	userRepository := repository.NewSqlcUserRepository(querier)
	gameRepository := repository.NewSqlcGameRepository(querier)

	hasher := security.NewHasher()

	authService := service.NewUserService(userRepository, hasher)
	gameService := service.NewGameService(gameRepository)

	authHandler := handler.NewAuthHandler(authService)
	gameHandler := handler.NewGameHandler(gameService)

	routes.SetupRoutes(g, authHandler, gameHandler)

	log.Fatal(g.Run(fmt.Sprintf("%s:%s", config.Api.Host, config.Api.Port)))

	logger.Log.Info("starting application")
}
