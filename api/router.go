package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/scelefeed/controller"
	"github.com/mustafasegf/scelefeed/repo"
	"github.com/mustafasegf/scelefeed/service"
)

func (s *Server) SetupRouter() {
	sceleRepo := repo.NewSceleRepo(s.Db)
	sceleService := service.NewSceleService(sceleRepo)
	sceleController := controller.NewSceleController(sceleService)

	lineService := service.NewLineService()
	lineController := controller.NewLineController(lineService, s.bot)
	s.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	s.router.POST("/login", sceleController.Login)
	s.router.POST("/callback", gin.WrapF(lineController.LineCallback))
}
