package api

import (
	"github.com/mustafasegf/scelefeed/controller"
	"github.com/mustafasegf/scelefeed/repo"
	"github.com/mustafasegf/scelefeed/service"
)

func (s *Server) SetupRouter() {
	sceleRepo := repo.NewSceleRepo(s.Db)
	sceleService := service.NewSceleService(sceleRepo)
	sceleController := controller.NewSceleController(sceleService)
	s.router.POST("/login", sceleController.Login)
}
