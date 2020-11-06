package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/registration"
	authMemory "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Staring server")

	userStore := memstore.New()
	userService := authMemory.NewService(userStore)
	registartionHandler := registration.NewHandler(userService)

	handler := chi.NewRouter()
	handler.Mount("/auth/register", registartionHandler)

	api := api.NewAPI(api.NewAPI(handler))
	srv := &http.Server{Addr: ":8080", Handler: api}
	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
