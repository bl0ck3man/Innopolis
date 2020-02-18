package repository

import (
	"context"
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/jmoiron/sqlx"
)

type (
	//go:generate mockery -name=user
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
	var user entities.User
	err := r.db.GetContext(r.ctx, &user, `INSERT INTO demo.public.users (name, email) VALUES  ($1, $2) RETURNING *`, name, email)

	return user, err
}

func (r *repo) Read(id int64) (entities.User, error) {
	var user entities.User
	err := r.db.GetContext(r.ctx, &user, `SELECT * FROM demo.public.users where id = $1`, id)

	return user, err
}

func (r *repo) UpdateEmail(id int64, email string) (entities.User, error) {
	var user entities.User
	err := r.db.GetContext(r.ctx, &user,
		`UPDATE demo.public.users set email = $2 where id = $1 RETURNING *;`, id, email)

	return user, err
}

func (r *repo) Delete(id int64) error {
	_, err := r.db.ExecContext(r.ctx, `DELETE FROM demo.public.users where id = $1`, id)

	return err
}