package main

import (
	"context"
	"flag"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"os"
	mongodb "server/internal/client"
	"time"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config      config
	logger      *log.Logger
	mongoClient *mongo.Database
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	mongoDBClient, err := mongodb.NewClient(context.Background(), "user-service")
	if err != nil {
		panic(err)
	}

	app := &application{
		config:      cfg,
		logger:      logger,
		mongoClient: mongoDBClient,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err = srv.ListenAndServe()
	logger.Fatal(err)

}
