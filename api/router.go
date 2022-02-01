package api

import (
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
	s.router.LoadHTMLGlob("templates/*")
	s.router.GET("/", sceleController.IndexPage)
	s.router.GET("/login", sceleController.LoginPage)
	s.router.POST("/login", sceleController.Login)
	s.router.POST("/callback/line", gin.WrapF(lineController.LineCallback))

}
