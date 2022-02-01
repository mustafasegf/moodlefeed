package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/scelefeed/line"
	"github.com/mustafasegf/scelefeed/scele"
)

func (s *Server) SetupRouter() {
	sceleRepo := scele.NewRepo(s.Db)
	sceleService := scele.NewService(sceleRepo)
	sceleController := scele.NewController(sceleService)

	schedule := scele.NewSchedule(sceleService, s.line)
	go schedule.RunSchedule()

	lineRepo := line.NewRepo(s.Db)
	lineService := line.NewService(lineRepo)
	lineController := line.NewController(s.line, lineService, sceleService)
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "wrong page. try to /login on line"})
	})
	s.router.LoadHTMLGlob("templates/*")
	s.router.GET("/login", sceleController.Index)
	s.router.POST("/login", sceleController.Login)
	s.router.POST("/callback/line", gin.WrapF(lineController.LineCallback))

}
