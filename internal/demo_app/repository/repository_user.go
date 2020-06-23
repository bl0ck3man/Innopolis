package repository

import (
	"context"
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/jmoiron/sqlx"
)

//go:generate mockery -name=User
type (
	User interface {
		Create(ctx context.Context, name, email string) (entities.User, error)
		Read(ctx context.Context,id int64) (entities.User, error)
		UpdateEmail(ctx context.Context, id int64, email string) (entities.User, error)
		Delete(ctx context.Context, id int64) error
	}

	repo struct {
		db *sqlx.DB
	}
)

func New(db *sqlx.DB) *repo {
	return &repo{
		db:  db,
	}
}

func (r *repo) Create(ctx context.Context, name, email string) (entities.User, error) {
	var user entities.User
	err := r.db.GetContext(ctx, &user, `INSERT INTO demo.public.users (name, email) VALUES  ($1, $2) RETURNING *`, name, email)

	return user, err
}

func (r *repo) Read(ctx context.Context, id int64) (entities.User, error) {
	var user entities.User
	err := r.db.GetContext(ctx, &user, `SELECT * FROM demo.public.users where id = $1`, id)

	return user, err
}

func (r *repo) UpdateEmail(ctx context.Context, id int64, email string) (entities.User, error) {
	var user entities.User
	err := r.db.GetContext(ctx, &user,
		`UPDATE demo.public.users set email = $2 where id = $1 RETURNING *;`, id, email)

	return user, err
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	_, err := r.db.ExecContext(ctx, `DELETE FROM demo.public.users where id = $1`, id)

	return err
}