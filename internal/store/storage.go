package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("records not found")
)

type Storage struct {
	Posts interface {
		GetById(context.Context, int64) (*Post, error)
		Create(context.Context, *Post) error
	}
	Users interface {
		Create(context.Context, *User) error
	}
}

func NewPostgresStoragedb(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &Userstore{db},
	}

}
