package repository

import (
	"study-gin/internal/models"
)

type UserRepository interface {
	Create(user *models.User) error
	FindByUuid(user *models.User, id int) error
}
