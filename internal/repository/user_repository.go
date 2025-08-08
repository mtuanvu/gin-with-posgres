package repository

import (
	"log"
	"study-gin/internal/models"
)

type SQLUserRepository struct {
}

func NewSQLUserRepository() *SQLUserRepository {
	return &SQLUserRepository{}
}

func (ur *SQLUserRepository) Create(user models.User) {
	log.Println("Create")
}

func (ur *SQLUserRepository) FindByUuid(id int) {
	log.Println("Find By UUID")
}
