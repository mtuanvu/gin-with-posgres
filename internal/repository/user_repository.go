package repository

import "log"

type SQLUserRepository struct {
}

func NewSQLUserRepository() *SQLUserRepository {
	return &SQLUserRepository{}
}

func (ur *SQLUserRepository) Create() {
	log.Println("Create")
}

func (ur *SQLUserRepository) FindByUuid() {
	log.Println("Find By UUID")
}