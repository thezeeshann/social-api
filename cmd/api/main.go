package main

import (
	"log"

	"github.com/thezeeshann/social/internal/db"
	"github.com/thezeeshann/social/internal/store"
)

const version = "0.0.1"

func main() {

	cfg := config{
		addr: ":8080",
		db: dbConfig{
			addr:         "postgresql://postgres:root@localhost:5432/social?sslmode=disable", // 5051 pg port
			maxOpenConns: 30,
			maxIdelConns: 30,
			maxIdleTime:  "15m",
		},
	}

	database, err := db.New(
		cfg.db.addr,
		cfg.db.maxIdelConns,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleTime,
	)

	store := store.NewPostgresStoragedb(database)

	app := &application{
		config: cfg,
		store:  store,
	}

	if err != nil {
		log.Panic(err)
	}

	defer database.Close()
	log.Printf("DB connected")

	mux := app.mount()
	log.Fatal(app.run(mux))
}
