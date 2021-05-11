package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/dwethmar/atami/pkg/api"
	"github.com/dwethmar/atami/pkg/api/handler"
	"github.com/dwethmar/atami/pkg/api/handler/beta"
	"github.com/dwethmar/atami/pkg/auth"
	"github.com/dwethmar/atami/pkg/config"
	"github.com/dwethmar/atami/pkg/database"
	"github.com/dwethmar/atami/pkg/domain"
	"github.com/dwethmar/atami/pkg/memstore"

	"github.com/go-chi/chi"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Staring server")

	var inMemory bool
	flag.BoolVar(&inMemory, "in-memory", false, "run atami with in memory storage")
	flag.Parse()

	c := config.Load()
	if err := c.Valid(); err != nil {
		fmt.Println("Not all env vars are set. Loading .env file.")
		c = config.LoadEnvFile()
		die(c.Valid())
	}

	var store *domain.Store

	if inMemory {
		fmt.Println("running server in-memory mode")
		store = domain.NewInMemoryStore(memstore.NewStore())
	} else {
		dataSource := database.GetPostgresDataSource(&database.PostgresConnectionConfig{
			DBHost:     c.DBHost,
			DBPort:     c.DBPort,
			DBUser:     c.DBUser,
			DBPassword: c.DBPassword,
			DBName:     "postgres",
		})
		db, err := database.Connect(c.DBDriverName, dataSource)
		die(err)
		defer db.Close()

		err = database.RunMigrations(db, c.DBName, c.MigrationFiles, c.DBMigrationVersion)
		die(err)

		store = domain.NewStore(db)
	}

	authService := auth.NewService(store.User.Finder)

	router := chi.NewRouter()
	router.Mount("/auth", handler.NewAuthRouter(authService, store))
	router.Mount("/beta/messages", beta.NewMessageRouter(store))

	srv, err := api.NewServer(":8080", router)
	die(err)
	srv.Start()
}

func die(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
