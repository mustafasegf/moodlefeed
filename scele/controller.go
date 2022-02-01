package scele

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/scelefeed/entity"
)

type Controller struct {
	svc *Service
}

func NewController(svc *Service) *Controller {
	return &Controller{
		svc: svc,
	}
}

func (ctrl *Controller) IndexPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", "")
}

func (ctrl *Controller) LoginPage(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", "")
}

func (ctrl *Controller) Login(ctx *gin.Context) {
	req := entity.LoginRequest{}
	err := ctx.BindJSON(&req)
	httprequest := HttpRequest{}
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
	token, err := httprequest.LoginScele(req.Username, req.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// get id
	sceleUser, err := httprequest.GetSceleId(token.Token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else if sceleUser.SceleID == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		return
	}
	// save user
	err = ctrl.svc.CreateUser(token.Token, lineID, sceleUser.SceleID)
	if err != nil {
		if strings.Contains(err.Error(), "unique") {
			err = errors.New("user already login")
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// get all course
	courses, err := httprequest.GetCourses(token.Token, sceleUser.SceleID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// save to db
	for _, course := range courses {
		_, err := ctrl.svc.CreateNewCourse(token.Token, sceleUser.SceleID, course)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	ctx.JSON(200, gin.H{"message": "good"})
}

func (ctrl *Controller) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
