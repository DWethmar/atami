package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/feed"
	"github.com/dwethmar/atami/pkg/api/login"
	"github.com/dwethmar/atami/pkg/api/registration"
	"github.com/dwethmar/atami/pkg/api/token"
	"github.com/dwethmar/atami/pkg/config"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/service"

	"github.com/go-chi/chi"
)

func main() {
	fmt.Println("Staring server")

	c := config.LoadEnvFile()
	dataSource := database.GetPostgresConnectionString(c)

	db, err := database.Connect(c.DBDriverName, dataSource)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := database.RunMigrations(db, c.DBName, c.MigrationFiles, c.DBMigrationVersion); err != nil {
		panic(err)
	}

	if secret, err := token.GetAccessSecret(); secret == nil || err != nil {
		panic(err)
	}

	authService := service.NewAuthServicePostgres(db)

	handler := chi.NewRouter()
	handler.Mount("/auth/register", registration.NewHandler(authService))
	handler.Mount("/auth/login", login.NewHandler(authService))
	handler.Mount("/feed", feed.NewHandler(authService))

	api := api.NewAPI(api.NewAPI(handler))
	srv := &http.Server{Addr: ":8080", Handler: api}
	defer srv.Close()

	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}
