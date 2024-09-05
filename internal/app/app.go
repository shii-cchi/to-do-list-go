package app

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"to-do-list-go/internal/config"
	"to-do-list-go/internal/database"
	"to-do-list-go/internal/delivery/handlers"
	"to-do-list-go/internal/service"
)

const (
	errLoadingConfig  = "error loading config"
	errConnectingToDB = "error connecting to db"

	successfulConfigLoad   = "config has been loaded successfully"
	successfulDBConnection = "successful connection to db"
	serverStart            = "server starting on port"
)

func Run() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf(errLoadingConfig+": %s\n", err)
	}
	log.Println(successfulConfigLoad)

	conn, err := sql.Open("postgres", fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName))
	if err != nil {
		log.Fatalf(errConnectingToDB+": %s\n", err)
	}
	log.Println(successfulDBConnection)
	repo := database.New(conn)

	s := service.NewService(repo)

	r := chi.NewRouter()
	h := handlers.NewHandler(s)
	h.RegisterRoutes(r)

	log.Printf(serverStart+" %s", cfg.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Port, r))
}
