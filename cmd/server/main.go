package main

import (
	"fmt"
	"log"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/api/handler/beta"
	"github.com/dwethmar/atami/pkg/config"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/service"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Staring server")

	c := config.Load()
	if err := c.Valid(); err != nil {
		fmt.Println("Not all env vars are set. Loading .env file.")
		c = config.LoadEnvFile()
		die(c.Valid())
	}

	dataSource := database.GetPostgresConnectionString(c)

	db, err := database.Connect(c.DBDriverName, dataSource)
	die(err)
	defer db.Close()

	err = database.RunMigrations(db, c.DBName, c.MigrationFiles, c.DBMigrationVersion)
	die(err)

	store := domain.NewStore(db);
	authService := service.NewAuthService(store.User.Finder, store.User.Creator)
	messageService := service.NewMessageServicePostgres(db)
	userService := service.NewUserServicePostgres(db)

	router := chi.NewRouter()
	router.Mount("/auth", handler.NewAuthRouter(authService, userService))
	router.Mount("/beta/messages", beta.NewMessageRouter(userService, messageService))

	srv, err := api.NewServer(":8081", router)
	die(err)
	srv.Start()
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
