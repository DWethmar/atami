package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/login"
	"github.com/dwethmar/atami/pkg/api/thread"
	"github.com/dwethmar/atami/pkg/api/token"

	"github.com/dwethmar/atami/pkg/api/registration"
	authMemory "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Staring server")

	if secret, err := token.GetAccessSecret(); secret == nil || err != nil {
		panic(err)
	}

	userStore := memstore.New()
	userService := authMemory.NewService(userStore)

	handler := chi.NewRouter()
	handler.Mount("/auth/register", registration.NewHandler(userService))
	handler.Mount("/auth/login", login.NewHandler(userService))

	handler.Mount("/threads", thread.NewHandler(userService))

	api := api.NewAPI(api.NewAPI(handler))
	srv := &http.Server{Addr: ":8081", Handler: api}
	log.Printf("Serving on :8081")
	log.Fatal(srv.ListenAndServe())
}
