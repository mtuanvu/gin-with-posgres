package handlers

import (
	"errors"
	"net/http"
	"strconv"
	"study-gin/internal/models"
	"study-gin/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
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
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid User Id",
		})

		return
	}

	var user models.User

	if err := uh.repo.FindByUuid(&user, id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}

func (uh *UserHandler) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	if err := uh.repo.Create(&user); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			ctx.JSON(http.StatusConflict, gin.H{
				"error": "email already exist",
			})

			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"data": user,
	})
}
