package main

import (
	"fmt"
	"log"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/beta/router"
	"github.com/dwethmar/atami/pkg/config"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/service"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Staring server")

	c := config.LoadEnvFile()
	die(c.Valid())

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
	handler.Mount("/beta/auth", router.NewAuthRouter(authService, userService))
	handler.Mount("/beta/messages", router.NewMessageRouter(userService, messageService))

	srv, err := api.NewServer(":8080", handler)
	die(err)
	srv.Start()
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
