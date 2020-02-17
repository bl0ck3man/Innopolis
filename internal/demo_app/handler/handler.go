package handler

import (
	"fmt"
	"net/http"
)

type (
	Handler interface {
		AddUser(w http.ResponseWriter, r *http.Request)
		GetUser(w http.ResponseWriter, r *http.Request)
		EditUser(w http.ResponseWriter, r *http.Request)
		RemoveUser(w http.ResponseWriter, r *http.Request)
	}

	handler struct{}
)

func New() *handler {
	return &handler{}
}

func (h *handler) AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
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
