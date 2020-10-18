package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/router"
)

func main() {
	fmt.Println("Staring server")
	api := api.NewAPI(router.NewRouter())
	srv := &http.Server{Addr: ":8080", Handler: api}
	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
