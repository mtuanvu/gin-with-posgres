package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"study-gin/internal/db/sqlc"
	"study-gin/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserResponse struct {
	UserID    int32     `json:"user_id"`
	Uuid      uuid.UUID `json:"uuid"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt string    `json:"created_at"`
}

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (uh *UserHandler) GetUserByUuid(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")
	userUUID, err := uuid.Parse(uuidParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User uuid",
		})

		return
	}

	user, err := uh.repo.FindByUuid(ctx, userUUID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := UserResponse{
		UserID:    user.UserID,
		Uuid:      user.Uuid,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02"),
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var input sqlc.CreateUserParams
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	userResponse, err := uh.repo.Create(ctx, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	response := UserResponse{
		UserID:    userResponse.UserID,
		Uuid:      userResponse.Uuid,
		Name:      userResponse.Name,
		Email:     userResponse.Email,
		CreatedAt: userResponse.CreatedAt.Format("2006-01-02"),
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": response,
	})
}
