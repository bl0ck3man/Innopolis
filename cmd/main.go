package main

import (
	"context"
	"github.com/blac3kman/Innopolis/internal/demo_app/handler"
	"github.com/blac3kman/Innopolis/internal/demo_app/repository"
	usecase_user "github.com/blac3kman/Innopolis/internal/demo_app/usecase"
	postgres "github.com/blac3kman/Innopolis/pkg"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	db := postgres.Connection
	repo := repository.New(context.Background(), db)
	useCase := usecase_user.New(repo)
	h := handler.New(useCase)

	router.HandleFunc("/user", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/add", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/user/edit", h.EditUser).Methods(http.MethodPost)
	router.HandleFunc("/user/delete", h.RemoveUser).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err.Error())
	}
}
