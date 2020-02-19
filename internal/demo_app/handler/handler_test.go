package handler_test

import (
	"github.com/blac3kman/Innopolis/internal/demo_app/handler"
	usecase_user "github.com/blac3kman/Innopolis/internal/demo_app/usecase"
	"github.com/blac3kman/Innopolis/internal/demo_app/usecase/mocks"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		us usecase_user.User
	}
	tests := []struct {
		name string
		args args
		want handler.Handler
	}{
		{
			name: `Success`,
			args: args{us: &mocks.User{}},
			want: handler.New(&mocks.User{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handler.New(tt.args.us); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*func Test_handler_AddUser(t *testing.T) {
	type fields struct {
		us usecase_user.User
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				us: tt.fields.us,
			}
		})
	}
}

func Test_handler_EditUser(t *testing.T) {
	type fields struct {
		us usecase_user.User
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				us: tt.fields.us,
			}
		})
	}
}

func Test_handler_GetUser(t *testing.T) {
	type fields struct {
		us usecase_user.User
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				us: tt.fields.us,
			}
		})
	}
}

func Test_handler_RemoveUser(t *testing.T) {
	type fields struct {
		us usecase_user.User
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &handler{
				us: tt.fields.us,
			}
		})
	}
}*/
