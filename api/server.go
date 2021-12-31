package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	Db     *gorm.DB
	line   *linebot.Client
}

func MakeServer(db *gorm.DB, line *linebot.Client) Server {
	router := gin.Default()
	server := Server{
		Db:     db,
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
