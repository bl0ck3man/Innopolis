package usecase_user

import (
	"context"

	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
)
//go:generate mockery -name=User
type (
	User interface {
		Get(ctx context.Context, id int64) (entities.User, error)
		Create(ctx context.Context, name, email string) (entities.User, error)
		UpdateEmail(ctx context.Context, id int64, email string) (entities.User, error)
		Delete(ctx context.Context, id int64) error
	}
	usecase struct {
		repo repository.User
	}
)

func New(repo repository.User) *usecase {
	return &usecase{
		repo: repo,
	}
}

func (u *usecase) Get(ctx context.Context, id int64) (entities.User, error) {
	return u.repo.Read(ctx, id)
}

func (u *usecase) Create(ctx context.Context, name, email string) (entities.User, error) {
	return u.repo.Create(ctx, name, email)
}

func (u *usecase) UpdateEmail(ctx context.Context, id int64, email string) (entities.User, error) {
	return u.repo.UpdateEmail(ctx, id, email)
}

func (u *usecase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}