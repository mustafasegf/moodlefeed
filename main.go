package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/mustafasegf/scelefeed/api"
	"github.com/mustafasegf/scelefeed/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:27017/", os.Getenv("MONGO_INITDB_ROOT_USERNAME"), os.Getenv("MONGO_INITDB_ROOT_PASSWORD"), os.Getenv("MONGO_HOST"))
	db, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Panic(err)
	}

	util.SetLogger()
	server := api.MakeServer(db)
	server.SetupRouter()
	server.RunServer(os.Getenv("SERVER_PORT"))

	defer func() {
		if err = db.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
}
