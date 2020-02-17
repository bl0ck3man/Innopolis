package usecase

import "github.com/blac3kman/Innopolis/internal/demo_app/entities"

type (
	//mockery -name=User
	User interface {
		Get(id int64) (entities.User, error)
		Create(name, email string) (entities.User, error)
		UpdateEmail(id int64, email string) (entities.User, error)
		Delete(id int64) error
	}
	usecase struct {}
)

func New() *usecase {
	return &usecase{}
}

func (u usecase) Get(id int64) (entities.User, error) {
	return entities.User{}, nil
}

func (u usecase) Create(name, email string) (entities.User, error) {
	return entities.User{}, nil
}

func (u usecase) UpdateEmail(id int64, email string) (entities.User, error) {
	return entities.User{}, nil
}

func (u usecase) Delete(id int64) error {
	return nil
}