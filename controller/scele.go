package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/service"
	"github.com/mustafasegf/scelefeed/util"
)

type Scele struct {
	svc *service.Scele
}

func NewSceleController(svc *service.Scele) *Scele {
	return &Scele{
		svc: svc,
	}
}

func (ctrl *Scele) Login(ctx *gin.Context) {
	req := entity.LoginRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	lineID := ctx.Query("id")
	if lineID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "no line id found"})
		return
	}
	token, err := util.LoginScele(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	ctrl.svc.SaveToken(token.Token, lineID)
	ctx.JSON(200, gin.H{"message": "good"})
}

func (ctrl *Scele) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
