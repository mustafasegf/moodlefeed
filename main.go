package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/scelefeed/api"
	"github.com/mustafasegf/scelefeed/core"
	"github.com/mustafasegf/scelefeed/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	util.SetLogger()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}))
	if err != nil {
		log.Fatal("canot load db: ", err)
	}

	bot, err := linebot.New(os.Getenv("LINE_SECRET"), os.Getenv("LINE_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	go core.RunSchedule()
	server := api.MakeServer(db, bot)
	server.RunServer()
}
