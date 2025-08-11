package repository

import (
	"context"
	"study-gin/internal/db/sqlc"
)

type UserRepository interface {
	Create(ctx context.Context, input sqlc.CreateUserParams) (sqlc.User, error)
	FindByUuid(id int)
}
