package main

import (
	"context"
	"flag"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	mongodb "server/internal/client"
	"server/internal/jsonlog"
	"server/internal/mailer"
	"sync"
)

var version = "1.0.0"

type config struct {
	port int
	env  string
	smtp struct {
		host     string
		port     int
		username string
		password string
		sender   string
	}
	limiter struct {
		rps     float64
		burst   int
		enabled bool
	}
}

type application struct {
	config      config
	logger      *jsonlog.Logger
	mongoClient *mongo.Database
	mailer      mailer.Mailer
	wg          sync.WaitGroup
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.smtp.host, "smtp-host", "smtp.gmail.com", "SMTP host")
	flag.IntVar(&cfg.smtp.port, "smtp-port", 465, "SMTP port")
	flag.StringVar(&cfg.smtp.username, "smtp-username", "healthyswitcher@gmail.com", "SMTP username")
	flag.StringVar(&cfg.smtp.password, "smtp-password", "zrlbtsggeqirgnlh", "SMTP password")
	flag.StringVar(&cfg.smtp.sender, "smtp-sender", "Healty-Swithcher <no-reply@healthy-switcher.olzhas.net>", "SMTP sender")

	flag.Float64Var(&cfg.limiter.rps, "limiter-rps", 2, "Rate limiter maximum requests per second")
	flag.IntVar(&cfg.limiter.burst, "limiter-burst", 4, "Rate limiter maximum burst")
	flag.BoolVar(&cfg.limiter.enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	mongoDBClient, err := mongodb.NewClient(context.Background(), "healthy-switcher")
	if err != nil {
		panic(err)
	}

	app := &application{
		config:      cfg,
		logger:      logger,
		mongoClient: mongoDBClient,
		mailer:      mailer.New(cfg.smtp.host, cfg.smtp.port, cfg.smtp.username, cfg.smtp.password, cfg.smtp.sender),
	}

	err = app.serve()
	if err != nil {
		logger.PrintFatal(err, nil)
	}

}
