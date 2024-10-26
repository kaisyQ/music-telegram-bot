package database

import (
	"context"
	"kaisyq/tg/music/core/env"
	"log"
	"sync"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	Pool *pgxpool.Pool
}

var (
	DatabaseInstance *Database
	databaseOnce     sync.Once
)

func GetInstance(ctx context.Context) (*Database, error) {

	env, err := env.GetInstance()

	if err != nil {
		log.Fatal("There was an error while trying to get env instance")
	}

	databaseOnce.Do(func() {
		db, err := pgxpool.New(ctx, env.DatabaseUrl)

		if err != nil {
			log.Fatalln("There was an error while trying to create database connection pool")
		}

		DatabaseInstance = &Database{db}
	})

	return DatabaseInstance, nil
}

func (database *Database) Close() {
	database.Pool.Close()
}
