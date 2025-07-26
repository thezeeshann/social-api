package store

import (
	"context"
	"database/sql"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

type Userstore struct {
	db *sql.DB
}

func (s *Userstore) Create(ctx context.Context, user *User) error {

	query := `INSERT INTO users (username, password, email) VALUES($1, $2, $3) RETURNING  id,created_at,updated_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Email,
		user.Password,
	).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdateAt,
	)

	if err != nil {
		return err
	}

	return nil
}
