package repository_test

import (
	"context"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
	"github.com/jmoiron/sqlx"
	"reflect"
	"testing"
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

/*func Test_repo_Create(t *testing.T) {
	type fields struct {
		ctx context.Context
		db  *sqlx.DB
	}
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
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
	type fields struct {
		ctx context.Context
		db  *sqlx.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
			if err := r.Delete(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_Read(t *testing.T) {
	type fields struct {
		ctx context.Context
		db  *sqlx.DB
	}
	type args struct {
		id int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
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
	type fields struct {
		ctx context.Context
		db  *sqlx.DB
	}
	type args struct {
		id    int64
		email string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entities.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &repo{
				ctx: tt.fields.ctx,
				db:  tt.fields.db,
			}
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
}*/
