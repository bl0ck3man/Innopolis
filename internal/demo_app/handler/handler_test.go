package handler_test

import (
	"errors"
	"github.com/blac3kman/Innopolis/internal/demo_app/entities"
	"github.com/blac3kman/Innopolis/internal/demo_app/handler"
	usecase_user "github.com/blac3kman/Innopolis/internal/demo_app/usecase"
	"github.com/blac3kman/Innopolis/internal/demo_app/usecase/mocks"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

type args struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func setUpArgs(method string, route string, payload string) args {
	req, err := http.NewRequest(method, route, strings.NewReader(payload))
	if err != nil {
		log.Fatal(err.Error())
	}

	return args{
		w: httptest.NewRecorder(),
		r: req,
	}
}

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

func Test_handler_AddUser(t *testing.T) {
	route := `/new/user`

	type fields struct {
		us usecase_user.User
	}
	tests := []struct {
		name           string
		fields         func() fields
		args           args
		wantStatusCode int
		wantBody       string
	}{
		{
			name: `Success`,
			fields: func() fields {

				mock := mocks.User{}
				mock.On(`Create`, `gopher`, `gopher@kazan.ru`).Return(entities.User{
					ID:    1,
					Name:  "gopher",
					Email: "gopher@kazan.ru",
				}, nil)

				return fields{
					us: &mock,
				}
			},
			args:           setUpArgs(http.MethodPost, route, `{"name": "gopher", "email": "gopher@kazan.ru"}`),
			wantStatusCode: http.StatusOK,
			wantBody:       `{"id":1,"name":"gopher","email":"gopher@kazan.ru"}`,
		},
		{
			name: `Bad request - empty email`,
			fields: func() fields {

				mock := mocks.User{}
				mock.On(`Create`, `gopher`, `gopher@kazan.ru`).Return(entities.User{
					ID:    1,
					Name:  "gopher",
					Email: "gopher@kazan.ru",
				}, nil)

				return fields{
					us: &mock,
				}
			},
			args:           setUpArgs(http.MethodPost, route, `{"name": "gopher"}`),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       http.StatusText(http.StatusBadRequest),
		},
		{
			name: `Bad request - empty name`,
			fields: func() fields {

				mock := mocks.User{}
				mock.On(`Create`, `gopher`, `gopher@kazan.ru`).Return(entities.User{
					ID:    1,
					Name:  "gopher",
					Email: "gopher@kazan.ru",
				}, nil)

				return fields{
					us: &mock,
				}
			},
			args:           setUpArgs(http.MethodPost, route, `{"email": "gopher@kazan.ru"}`),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       http.StatusText(http.StatusBadRequest),
		},
		{
			name: `Bad request - empty name`,
			fields: func() fields {

				mock := mocks.User{}
				mock.On(`Create`, `gopher`, `gopher@kazan.ru`).Return(entities.User{}, errors.New(`some error`))

				return fields{
					us: &mock,
				}
			},
			args:           setUpArgs(http.MethodPost, route, `{"name": "gopher", "email": "gopher@kazan.ru"}`),
			wantStatusCode: http.StatusInternalServerError,
			wantBody:       http.StatusText(http.StatusInternalServerError),
		},
		{
			name: `Bad request - empty payload`,
			fields: func() fields {

				mock := mocks.User{}

				return fields{
					us: &mock,
				}
			},
			args:           setUpArgs(http.MethodPost, route, ``),
			wantStatusCode: http.StatusBadRequest,
			wantBody:       http.StatusText(http.StatusBadRequest),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := http.HandlerFunc(handler.New(tt.fields().us).AddUser)
			h.ServeHTTP(tt.args.w, tt.args.r)

			gotBody := strings.TrimSpace(tt.args.w.Body.String())

			assert.Equal(t, tt.wantStatusCode, tt.args.w.Code)
			assert.Equal(t, tt.wantBody, gotBody)
		})
	}
}

/*func Test_handler_EditUser(t *testing.T) {
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
