package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/m-bromo/my-game-list/internal/web/handler"
)

func SetupRoutes(g *gin.Engine, ah *handler.AuthHandler, gh *handler.GameHandler) {
	g.POST("/register", ah.Register)
	g.POST("/login", ah.Login)

	g.GET("/games", gh.GetAll)
}
