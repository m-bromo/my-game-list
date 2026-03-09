package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/m-bromo/my-game-list/internal/service"
	"github.com/m-bromo/my-game-list/internal/web/models"
)

type AuthHandler struct {
	authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (a *AuthHandler) Register(c *gin.Context) {
	var input models.RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	if err := a.authService.Register(c.Request.Context(), &input); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, "User created")
}

func (a AuthHandler) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(err)
		return
	}

	if err := a.authService.Login(c.Request.Context(), &input); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, "Logged in")
}
