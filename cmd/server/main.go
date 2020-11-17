package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/router"
	"github.com/dwethmar/atami/pkg/config"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/service"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Staring server")

	c := config.LoadEnvFile()
	err := c.Valid()
	die(err)

	dataSource := database.GetPostgresConnectionString(c)

	db, err := database.Connect(c.DBDriverName, dataSource)
	die(err)
	defer db.Close()

	err = database.RunMigrations(db, c.DBName, c.MigrationFiles, c.DBMigrationVersion)
	die(err)

	userService := service.NewUserServicePostgres(db)
	authService := service.NewAuthServicePostgres(db)
	messageService := service.NewMessageServicePostgres(db)

	handler := chi.NewRouter()
	handler.Mount("/auth", router.NewAuthRouter(authService, userService))
	handler.Mount("/messages", router.NewMessageRouter(userService, messageService))

	api := api.NewAPI(api.NewAPI(handler))
	srv := &http.Server{Addr: ":8080", Handler: api}
	defer srv.Close()

	log.Printf("Serving on :8080")
	log.Fatal(srv.ListenAndServe())
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
