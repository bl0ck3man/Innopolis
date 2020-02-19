package handler

import (
	"database/sql"
	"encoding/json"
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
	w.Header().Set(`Content-type`, `application/json`)

	type payload struct {
		Id int64 `json:"user_id"`
	}

	p := new(payload)

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if p.Id == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	user, err := h.us.Get(p.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *handler) EditUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-type`, `application/json`)

	type payload struct {
		Id int64 `json:"user_id"`
		Email string `json:"email"`
	}

	p := new(payload)

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if p.Id == 0 || len(p.Email) == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	user, err := h.us.UpdateEmail(p.Id, p.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	json.NewEncoder(w).Encode(user)
}

func (h *handler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(`Content-type`, `application/json`)

	type payload struct {
		Id int64 `json:"user_id"`
	}

	p := new(payload)

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	if p.Id == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)

		return
	}

	err := h.us.Delete(p.Id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)

			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
