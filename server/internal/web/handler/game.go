package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/my-game-list/internal/service"
)

type GameHandler struct {
	gameService service.GameService
}

func NewGameHandler(gameService service.GameService) *GameHandler {
	return &GameHandler{
		gameService: gameService,
	}
}

func (h *GameHandler) GetAll(c *gin.Context) {
	games, err := h.gameService.GetAll(c.Request.Context())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusNotFound, games)
}
