package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"study-gin/internal/models"
)

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(DB *sql.DB) *SQLUserRepository {
	return &SQLUserRepository{
		db: DB,
	}
}

func (ur *SQLUserRepository) Create(user models.User) error {
	row := ur.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING user_id", user.Name, user.Email)
	err := row.Scan(&user.Id)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (ur *SQLUserRepository) FindByUuid(id int) (*models.User, error) {
	var user models.User
	err := ur.db.QueryRow("SELECT * FROM users WHERE user_id = $1", id).Scan(&user.Id, &user.Name, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &user, nil
}
