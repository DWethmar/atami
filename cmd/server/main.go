package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	userApi "github.com/dwethmar/atami/pkg/api/user"
	authMemory "github.com/dwethmar/atami/pkg/auth/memory"
	"github.com/dwethmar/atami/pkg/memstore"
)

func main() {
	fmt.Println("Staring server")

	userStore := memstore.New()
	userService := authMemory.NewService(userStore)
	userHandler := userApi.NewHandler(userService)

	api := api.NewAPI(api.NewAPI(userHandler))
	srv := &http.Server{Addr: ":8080", Handler: api}
	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
