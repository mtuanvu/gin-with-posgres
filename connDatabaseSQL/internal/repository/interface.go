package repository

import "study-gin/internal/models"

type UserRepository interface {
	Create(user models.User) error
	FindByUuid(id int) (*models.User, error)
}
