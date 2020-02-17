package main

import (
	"github.com/blac3kman/Innopolis/internal/demo_app/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	h := handler.New()

	router.HandleFunc("/user", h.GetUser).Methods(http.MethodGet)
	router.HandleFunc("/user/add", h.AddUser).Methods(http.MethodPost)
	router.HandleFunc("/user/edit", h.EditUser).Methods(http.MethodPost)
	router.HandleFunc("/user/delete", h.RemoveUser).Methods(http.MethodPost)

	if err := http.ListenAndServe(":8080", router); err != nil {
		panic(err.Error())
	}
}
