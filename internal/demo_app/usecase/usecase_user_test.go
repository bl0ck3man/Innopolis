package usecase_user_test

import (
	"errors"
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository/mocks"
	usecase_user "github.com/blac3kman/Innopolis/internal/demo_app/usecase"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		repo repository.User
	}
	tests := []struct {
		name string
		args args
		want usecase_user.User
	}{
		{
			name: `Success`,
			args: args{&mocks.User{}},
			want: usecase_user.New(&mocks.User{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := usecase_user.New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_Create(t *testing.T) {
	type fields struct {
		repo repository.User
	}
	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		fields  func(args args, want entities.User) fields
		args    args
		want    entities.User
		wantErr bool
	}{
		{
			name: `Success`,
			fields: func(args args, want entities.User) fields {
				m := &mocks.User{}
				m.On(`Create`, args.name, args.email).Return(want, nil)

				return fields{
					repo: m,
				}
			},
			args: args{
				name:  "gopher",
				email: "gopner@innopolis.ru",
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `gopher@innopolis.ru`,
			},
			wantErr: false,
		},
		{
			name: `Error`,
			fields: func(args args, want entities.User) fields {
				m := &mocks.User{}
				m.On(`Create`, args.name, args.email).Return(want, errors.New(`some error`))

				return fields{
					repo: m,
				}
			},
			args: args{
				name:  "gopher",
				email: "gopner@innopolis.ru",
			},
			want:    entities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields(tt.args, tt.want).repo
			u := usecase_user.New(repo)

			got, err := u.Create(tt.args.name, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_Delete(t *testing.T) {
	type fields struct {
		repo repository.User
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  func(args args) fields
		args    args
		wantErr bool
	}{
		{
			name: `Success`,
			fields: func(args args) fields {
				m := &mocks.User{}
				m.On(`Delete`, args.id).Return(nil)

				return fields{
					repo: m,
				}
			},
			args: args{
				id: 1,
			},
			wantErr: false,
		},
		{
			name: `Error`,
			fields: func(args args) fields {
				m := &mocks.User{}
				m.On(`Delete`, args.id).Return(errors.New(`some error`))

				return fields{
					repo: m,
				}
			},
			args: args{
				id: 1,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields(tt.args).repo
			u := usecase_user.New(repo)

			if err := u.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_usecase_Get(t *testing.T) {
	type fields struct {
		repo repository.User
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  func(args args, want entities.User) fields
		args    args
		want    entities.User
		wantErr bool
	}{
		{
			name: `Success`,
			fields: func(args args, want entities.User) fields {
				m := &mocks.User{}
				m.On(`Read`, args.id).Return(want, nil)

				return fields{
					repo: m,
				}
			},
			args: args{
				id: 1,
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `gopher@innopolis.ru`,
			},
			wantErr: false,
		},
		{
			name: `Error`,
			fields: func(args args, want entities.User) fields {
				m := &mocks.User{}
				m.On(`Read`, args.id).Return(want, errors.New(`some error`))

				return fields{
					repo: m,
				}
			},
			args: args{
				id: 1,
			},
			want:    entities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields(tt.args, tt.want).repo
			u := usecase_user.New(repo)

			got, err := u.Get(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_usecase_UpdateEmail(t *testing.T) {
	type fields struct {
		repo repository.User
	}
	type args struct {
		id    int64
		email string
	}
	tests := []struct {
		name    string
		fields  func(args args, want entities.User) fields
		args    args
		want    entities.User
		wantErr bool
	}{
		{
			name: `Success`,
			fields: func(args args, want entities.User) fields {
				m := &mocks.User{}
				m.On(`UpdateEmail`, args.id, args.email).Return(want, nil)

				return fields{
					repo: m,
				}
			},
			args: args{
				id:    1,
				email: `newGopher@innopolis.ru`,
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `newGopher@innopolis.ru`,
			},
			wantErr: false,
		},
		{
			name: `Error`,
			fields: func(args args, want entities.User) fields {
				m := &mocks.User{}
				m.On(`UpdateEmail`, args.id, args.email).Return(want, errors.New(`some Error`))

				return fields{
					repo: m,
				}
			},
			args: args{
				id:    1,
				email: `newGopher@innopolis.ru`,
			},
			want:    entities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.fields(tt.args, tt.want).repo
			u := usecase_user.New(repo)

			got, err := u.UpdateEmail(tt.args.id, tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}
