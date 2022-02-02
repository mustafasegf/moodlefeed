package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Server struct {
	router *gin.Engine
	DB     *pgxpool.Pool
	line   *linebot.Client
}

func MakeServer(db *pgxpool.Pool, line *linebot.Client) Server {
	router := gin.Default()
	server := Server{
		DB:     db,
		router: router,
		line:   line,
	}
	return server
}

func (s *Server) RunServer() {
	s.SetupRouter()
	serverString := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	s.router.Run(serverString)
}
