package repository

import (
	"context"
	"fmt"
	"log"
	"study-gin/internal/db/sqlc"
)

type SQLUserRepository struct {
	db sqlc.Querier
}

func NewSQLUserRepository(DB sqlc.Querier) *SQLUserRepository {
	return &SQLUserRepository{
		db: DB,
	}
}

func (ur *SQLUserRepository) Create(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error) {
	user, err := ur.db.CreateUser(ctx, input)
	if err != nil {
		return sqlc.User{}, fmt.Errorf("failed to create user: %w", err)
	}

	return user, nil
}

func (ur *SQLUserRepository) FindByUuid(id int) {
	log.Println("Find By UUID")
}
