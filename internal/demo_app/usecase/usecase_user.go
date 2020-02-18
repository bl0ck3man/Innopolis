package usecase_user

import (
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
)

type (
	//mockery -name=User
	User interface {
		Get(id int64) (entities.User, error)
		Create(name, email string) (entities.User, error)
		UpdateEmail(id int64, email string) (entities.User, error)
		Delete(id int64) error
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

func (u *usecase) Get(id int64) (entities.User, error) {
	return u.repo.Read(id)
}

func (u *usecase) Create(name, email string) (entities.User, error) {
	return u.repo.Create(name, email)
}

func (u *usecase) UpdateEmail(id int64, email string) (entities.User, error) {
	return u.repo.UpdateEmail(id, email)
}

func (u *usecase) Delete(id int64) error {
	return u.repo.Delete(id)
}