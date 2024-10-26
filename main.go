package main

import (
	"context"
	"fmt"
	env "kaisyq/tg/music/core/env"
	"kaisyq/tg/music/infrastructure"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	var env, err = env.GetInstance()

	if err != nil {
		log.Fatalln("Error occures while trying to load environment variables", err.Error())
		panic(err.Error())
	}

	dbpool, err := pgxpool.New(context.Background(), env.DatabaseUrl)

	fmt.Println(env.DatabaseUrl)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}

	defer dbpool.Close()

	var greeting string

	err = dbpool.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)

	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	var telegramBot infrastructure.TelegramBot = infrastructure.TelegramBot{Token: env.TelegramBotToken}

	telegramBot.Init()
}
