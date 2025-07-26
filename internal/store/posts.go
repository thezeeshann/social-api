package store

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64    `json:"id"`
	Content   string   `json:"content"`
	Title     string   `json:"title"`
	UserId    int16    `json:"user_id"`
	Tags      []string `json:"tags"`
	CreatedAt string   `json:"created_at"`
	UpdateAt  string   `json:"updated_at"`
}

type PostsStore struct {
	db *sql.DB
}

func (s *PostsStore) Create(ctx context.Context, post *Post) error {

	query := `INSERT INTO posts (content, title, user_id, tags) VALUES($1, $2, $3, $4) RETURNING  id,created_at, updated_at`

	err := s.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserId,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdateAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostsStore) GetById(ctx context.Context, id int64) (*Post, error) {
	query := `SELECT id,title,user_id,content,created_at,updated_at FROM posts WHERE id=$1`
	var posts Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&posts.ID,
		&posts.Title,
		&posts.UserId,
		&posts.Content,
		&posts.CreatedAt,
		&posts.UpdateAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}
	return &posts, nil
}
