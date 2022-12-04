package main

import (
	"fmt"
	"net/http"

	"github.com/celiovjunior/goserverwidget/configs"
	"github.com/celiovjunior/goserverwidget/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	// configuring routers
	r := chi.NewRouter()
	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Get)

	fmt.Println("Server is running 🐎")
	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
