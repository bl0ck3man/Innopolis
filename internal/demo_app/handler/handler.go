package handler

import (
	"encoding/json"
	"fmt"
	usecase_user "github.com/blac3kman/Innopolis/internal/demo_app/usecase"
	"net/http"
)

type (
	Handler interface {
		AddUser(w http.ResponseWriter, r *http.Request)
		GetUser(w http.ResponseWriter, r *http.Request)
		EditUser(w http.ResponseWriter, r *http.Request)
		RemoveUser(w http.ResponseWriter, r *http.Request)
	}

	handler struct{
		us usecase_user.User
	}
)

func New(us usecase_user.User) *handler {
	return &handler{
		us: us,
	}
}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-type`, `application/json`)

	type payload struct {
		Name string `json:"name"`
		Email string `json:"email"`
	}

	p := new(payload)

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if len(p.Email) == 0 || len(p.Name) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	user, err := h.us.Create(p.Name, p.Email)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (h *handler) EditUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func (h *handler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
