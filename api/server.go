package api

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	router *gin.Engine
	Db     *gorm.DB
}

func MakeServer(db *gorm.DB) Server {
	router := gin.Default()
	server := Server{
		Db:     db,
		router: router,
	}
	return server
}

func (s *Server) RunServer() {
	s.SetupRouter()
	serverString := fmt.Sprintf(":%s", os.Getenv("SERVER_PORT"))
	s.router.Run(serverString)
}
