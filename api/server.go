package api

import (
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	Db *mongo.Client
}

func MakeServer(db *mongo.Client) Server {
	server := Server{
		Db: db,
	}
	return server
}

func (s *Server) RunServer(port string) {
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
