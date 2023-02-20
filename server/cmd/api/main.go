package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	mongodb "server/internal/client"
	"server/internal/data/user"
	"server/internal/data/user/db"
	"time"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	mongoDBClient, err := mongodb.NewClient(context.Background(), "user-service")
	if err != nil {
		panic(err)
	}
	storage := db.NewStorage(mongoDBClient, "users")

	user1 := user.User{

		Email:    "olzhasayato@gmail.com",
		Name:     "Olzhas",
		Password: "Test",
	}

	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}

	fmt.Println(user1ID)

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}
