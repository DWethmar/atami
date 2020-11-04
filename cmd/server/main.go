package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	userApi "github.com/dwethmar/atami/pkg/api/user"
	"github.com/dwethmar/atami/pkg/memstore"
	"github.com/dwethmar/atami/pkg/usecase/userusecase"
	userMemory "github.com/dwethmar/atami/pkg/user/memory"
)

func main() {
	fmt.Println("Staring server")

	userStore := memstore.New()
	userUsecase := userusecase.NewUserUsecase(userMemory.NewService(userStore))
	userHandler := userApi.NewHandler(userUsecase)

	api := api.NewAPI(api.NewAPI(userHandler))
	srv := &http.Server{Addr: ":8080", Handler: api}
	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
