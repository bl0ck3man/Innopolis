package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
)

type fields struct {
	ctx     context.Context
	sqlx    *sqlx.DB
	sqlmock sqlmock.Sqlmock
}

func getFieldsMocks() fields {
	db, dbMock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	xdb := sqlx.NewDb(db, `sqlmock`)

	return fields{
		ctx:     context.TODO(),
		sqlx:    xdb,
		sqlmock: dbMock,
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ctx context.Context
		db  *sqlx.DB
	}
	tests := []struct {
		name string
		args args
		want repository.User
	}{
		{
			name: "Success New user repository",
			args: args{
				ctx: context.TODO(),
				db:  &sqlx.DB{},
			},
			want: repository.New(context.TODO(), &sqlx.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.New(tt.args.ctx, tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_Create(t *testing.T) {
	queryTml := `INSERT INTO demo.public.users (name, email) VALUES  ($1, $2) RETURNING *`
	rowsTml := []string{`id`, `name`, `email`}

	type args struct {
		name  string
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.User
		wantErr bool
		setMock func(mock sqlmock.Sqlmock, args args, want entities.User)
	}{
		{
			name: "Success",
			fields: getFieldsMocks(),
			args:args{
				name:  "gopher",
				email: "gopher@innopolis.ru",
			},
			want: entities.User{
				ID:    1,
				Name:  "gopher",
				Email: "gopher@innopolis.ru",
			},
			wantErr: false,
			setMock: func(mock sqlmock.Sqlmock, args args, want entities.User) {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.name, args.email).
					WillReturnRows(rows)
			},
		},
		{
			name: "Error",
			fields: getFieldsMocks(),
			args: args{
				name:  "gopher",
				email: "gopher@innopolis.ru",
			},
			want: entities.User{},
			wantErr: true,
			setMock: func(mock sqlmock.Sqlmock, args args, want entities.User) {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.name, args.email).
					WillReturnError(errors.New(`some error`))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.New(tt.fields.ctx, tt.fields.sqlx)
			tt.setMock(tt.fields.sqlmock, tt.args, tt.want)

			got, err := r.Create(tt.args.name, tt.args.email)
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

func Test_repo_Delete(t *testing.T) {
	queryTml := "DELETE FROM demo.public.users where id = $1"

	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       error
		wantErr    bool
		setUpMocks func(mock sqlmock.Sqlmock, args args, want error)
	}{
		{
			name:   "Success delete user",
			fields: getFieldsMocks(),
			args: args{
				id: 1,
			},
			want:    nil,
			wantErr: false,
			setUpMocks: func(mock sqlmock.Sqlmock, args args, want error) {
				mock.ExpectExec(queryTml).
					WithArgs(args.id).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name:   "Error delete user",
			fields: getFieldsMocks(),
			args: args{
				id: 1,
			},
			want:    sql.ErrNoRows,
			wantErr: true,
			setUpMocks: func(mock sqlmock.Sqlmock, args args, want error) {
				mock.ExpectQuery(queryTml).
					WithArgs(args.id).
					WillReturnError(want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.New(tt.fields.ctx, tt.fields.sqlx)
			tt.setUpMocks(tt.fields.sqlmock, tt.args, tt.want)

			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_Read(t *testing.T) {
	rowsTml := []string{`id`, `name`, `email`}
	queryTml := "SELECT * FROM demo.public.users where id = $1"

	type args struct {
		id int64
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       entities.User
		wantErr    bool
		setUpMocks func(mock sqlmock.Sqlmock, args args, want entities.User)
	}{
		{
			name:   "Success read user",
			fields: getFieldsMocks(),
			args: args{
				id: 1,
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `gopher@innopolis.ru`,
			},
			wantErr: false,
			setUpMocks: func(mock sqlmock.Sqlmock, args args, want entities.User) {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.id).
					WillReturnRows(rows)
			},
		},
		{
			name:   "Error_read_user",
			fields: getFieldsMocks(),
			args: args{
				id: 1,
			},
			want:    entities.User{},
			wantErr: true,
			setUpMocks: func(mock sqlmock.Sqlmock, args args, want entities.User) {
				mock.ExpectQuery(queryTml).
					WithArgs(args.id).
					WillReturnError(sql.ErrNoRows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := repository.New(tt.fields.ctx, tt.fields.sqlx)
			tt.setUpMocks(tt.fields.sqlmock, tt.args, tt.want)

			got, err := r.Read(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Read() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Read() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_UpdateEmail(t *testing.T) {
	rowsTml := []string{`id`, `name`, `email`}
	queryTml := "UPDATE demo.public.users set email = $2 where id = $1 RETURNING *;"

	type args struct {
		id    int64
		email string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		want       entities.User
		wantErr    bool
		setUpMocks func(mock sqlmock.Sqlmock, args args, want entities.User)
	}{
		{
			name:   "Success update email user",
			fields: getFieldsMocks(),
			args: args{
				id:    1,
				email: `updatedGopher@innopolis.ru`,
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `updatedGopher@innopolis.ru`,
			},
			wantErr: false,
			setUpMocks: func(mock sqlmock.Sqlmock, args args, want entities.User) {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.id, args.email).
					WillReturnRows(rows)
			},
		},
		{
			name:   "Error update email user",
			fields: getFieldsMocks(),
			args: args{
				id:    1,
				email: `updatedGopher@innopolis.ru`,
			},
			want:    entities.User{},
			wantErr: true,
			setUpMocks: func(mock sqlmock.Sqlmock, args args, want entities.User) {
				mock.ExpectQuery(queryTml).
					WithArgs(args.id, args.email).
					WillReturnError(sql.ErrNoRows)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := repository.New(tt.fields.ctx, tt.fields.sqlx)
			tt.setUpMocks(tt.fields.sqlmock, tt.args, tt.want)

			got, err := r.UpdateEmail(tt.args.id, tt.args.email)
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
