package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/scelefeed/api"
	"github.com/mustafasegf/scelefeed/util"
)

func main() {

	godotenv.Load()
	util.SetLogger()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	ctx := context.Background()
	db, err := pgxpool.Connect(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		db.Close()
	}()

	bot, err := linebot.New(os.Getenv("LINE_SECRET"), os.Getenv("LINE_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	server := api.MakeServer(db, bot)
	server.RunServer()
}
