package repository_test

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"

	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
)

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
			want: repository.New(&sqlx.DB{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_Create(t *testing.T) {
	queryTml := `INSERT INTO demo.public.users (name, email) VALUES  ($1, $2) RETURNING *`
	rowsTml := []string{`id`, `name`, `email`}

	type fields struct {
		sqlx    *sqlx.DB
	}
	type args struct {
		ctx   context.Context
		name  string
		email string
	}
	tests := []struct {
		name    string
		fields  func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields
		args    args
		want    entities.User
		wantErr bool
		setMock func(mock sqlmock.Sqlmock, args args, want entities.User)
	}{
		{
			name:   "Success",
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.name, args.email).
					WillReturnRows(rows)

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			args: args{
				ctx: context.TODO(),
				name:  "gopher",
				email: "gopher@kaliningrad.ru",
			},
			want: entities.User{
				ID:    1,
				Name:  "gopher",
				Email: "gopher@kaliningrad.ru",
			},
			wantErr: false,
		},
		{
			name:   "Error",
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.name, args.email).
					WillReturnError(errors.New(`some error`))

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			args: args{
				ctx: context.TODO(),
				name:  "gopher",
				email: "gopher@kaliningrad.ru",
			},
			want:    entities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sqmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			fields := tt.fields(tt.args, db, sqmock, tt.want)
			r := repository.New(fields.sqlx)

			got, err := r.Create(tt.args.ctx, tt.args.name, tt.args.email)
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

	type fields struct {
		sqlx    *sqlx.DB
	}
	type args struct {
		ctx context.Context
		id int64
	}
	tests := []struct {
		name       string
		fields 	   func(args args, db *sql.DB, mock sqlmock.Sqlmock) fields
		args       args
		want       error
		wantErr    bool
	}{
		{
			name:   "Success_delete_user",
			args: args{
				ctx: context.TODO(),
				id: 1,
			},
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock) fields {
				mock.ExpectExec(queryTml).
					WithArgs(args.id).
					WillReturnResult(sqlmock.NewResult(1, 1))

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			want:    nil,
			wantErr: false,
		},
		{
			name:   "Error_delete_user",
			args: args{
				ctx: context.TODO(),
				id: 1,
			},
			want:    sql.ErrNoRows,
			wantErr: true,
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock) fields {
				mock.ExpectQuery(queryTml).
					WithArgs(args.id).
					WillReturnError(errors.New(`some error`))

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sqmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			fields := tt.fields(tt.args, db, sqmock)
			r := repository.New(fields.sqlx)

			if err := r.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_Read(t *testing.T) {
	rowsTml := []string{`id`, `name`, `email`}
	queryTml := "SELECT * FROM demo.public.users where id = $1"

	type fields struct {
		sqlx    *sqlx.DB
	}
	type args struct {
		ctx context.Context
		id int64
	}
	tests := []struct {
		name       string
		fields     func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields
		args       args
		want       entities.User
		wantErr    bool
	}{
		{
			name:   "Success_read_user",
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.id).
					WillReturnRows(rows)

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			args: args{
				ctx: context.TODO(),
				id: 1,
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `gopher@kaliningrad.ru`,
			},
			wantErr: false,
		},
		{
			name:   "Error_read_user",
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields {
				mock.ExpectQuery(queryTml).
					WithArgs(args.id).
					WillReturnError(sql.ErrNoRows)

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			args: args{
				ctx: context.TODO(),
				id: 1,
			},
			want:    entities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sqmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			fields := tt.fields(tt.args, db, sqmock, tt.want)
			r := repository.New(fields.sqlx)

			got, err := r.Read(tt.args.ctx, tt.args.id)
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

	type fields struct {
		sqlx    *sqlx.DB
	}
	type args struct {
		ctx   context.Context
		id    int64
		email string
	}
	tests := []struct {
		name       string
		fields     func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields
		args       args
		want       entities.User
		wantErr    bool
	}{
		{
			name:   "Success update email user",
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields {
				rows := sqlmock.NewRows(rowsTml)

				rows.AddRow(
					want.ID,
					want.Name,
					want.Email,
				)

				mock.ExpectQuery(queryTml).
					WithArgs(args.id, args.email).
					WillReturnRows(rows)

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			args: args{
				ctx: context.TODO(),
				id:    1,
				email: `updatedgopher@kaliningrad.ru`,
			},
			want: entities.User{
				ID:    1,
				Name:  `gopher`,
				Email: `updatedgopher@kaliningrad.ru`,
			},
			wantErr: false,
		},
		{
			name:   "Error update email user",
			fields: func(args args, db *sql.DB, mock sqlmock.Sqlmock, want entities.User) fields {
				mock.ExpectQuery(queryTml).
					WithArgs(args.id, args.email).
					WillReturnError(sql.ErrNoRows)

				return fields{
					sqlx: sqlx.NewDb(db, "sqlmock"),
				}
			},
			args: args{
				ctx: context.TODO(),
				id:    1,
				email: `updatedgopher@kaliningrad.ru`,
			},
			want:    entities.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, sqmock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			fields := tt.fields(tt.args, db, sqmock, tt.want)
			r := repository.New(fields.sqlx)

			got, err := r.UpdateEmail(tt.args.ctx, tt.args.id, tt.args.email)
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
