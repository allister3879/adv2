package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"greenlight/internal/data"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   string
}

type application struct {
	config config
	logger *log.Logger
	models data.DBModel
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")

	flag.StringVar(&cfg.db, "db", "postgres://username:password@localhost/s.alshynovDB?sslmode=disable", "Database connection string")

	flag.Parse()
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(cfg)
	if err != nil {
		logger.Fatalf("Connection failed. Error is: %s", err)
	}

	// db will be closed before main function is completed.
	defer db.Close()
	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db), // data.NewModels() function to initialize a Models struct
	}
	// Use the httprouter instance returned by app.routes() as the server handler.
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	// reuse defined variable err
	err = srv.ListenAndServe()
	logger.Fatal(err)
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	//context with a 5 second timeout deadline
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = db.PingContext(ctx) //create a connection and verify that everything is set up correctly.

	if err != nil {
		return nil, err
	}

	return db, nil
}
