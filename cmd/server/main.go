package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
)

func main() {
	fmt.Println("Staring server")
	api := api.NewAPI(api.NewAPI())
	srv := &http.Server{Addr: ":8080", Handler: api}
	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
