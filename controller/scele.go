package controller

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

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
	// login
	token, err := util.LoginScele(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id
	sceleUser, err := util.GetSceleId(token.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// save user
	err = ctrl.svc.CreateUser(token.Token, lineID, sceleUser.SceleID)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			err = errors.New("user already login")
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// get all course
	courses, err := util.GetCourses(token.Token, sceleUser.SceleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// save to db
	for _, course := range courses {
		courseDetail, err := ctrl.svc.CreateNewCourse(token.Token, sceleUser.SceleID, course)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		fmt.Printf("course: %#v\n", courseDetail)
	}

	ctx.JSON(200, gin.H{"message": "good"})
}

func (ctrl *Scele) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
