package repository_user

import (
	"context"
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/jmoiron/sqlx"
)

type (
	//mockery -name=user
	User interface {
		Create(name, email string) (entities.User, error)
		Read(id int64) (entities.User, error)
		UpdateEmail(id int64, email string) (entities.User, error)
		Delete(id int64) error
	}

	repo struct {
		ctx context.Context
		db *sqlx.DB
	}
)

func New(ctx context.Context, db *sqlx.DB) *repo {
	return &repo{
		ctx: ctx,
		db:  db,
	}
}

func (r *repo) Create(name, email string) (entities.User, error) {
	return entities.User{}, nil
}

func (r *repo) Read(id int64) (entities.User, error) {
	return entities.User{}, nil
}

func (r *repo) UpdateEmail(id int64, email string) (entities.User, error) {
	return entities.User{}, nil
}

func (r *repo) Delete(id int64) error {
	return nil
}