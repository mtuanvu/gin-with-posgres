package handlers

import (
	"net/http"
	"study-gin/internal/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	uh.repo.FindByUuid()
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Get User By UUID",
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	uh.repo.Create()
	ctx.JSON(http.StatusOK, gin.H{
		"data": "Create User",
	})
}

