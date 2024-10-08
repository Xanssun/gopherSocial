package store

import (
	"context"
	"database/sql"

	"github.com/Xanssun/gopherSocial.git/internal/models"
)

type Storage struct {
	Posts interface {
		Create(context.Context, *models.Post) error
	}

	Users interface {
		Create(context.Context) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}
