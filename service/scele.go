package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/repo"
	"github.com/mustafasegf/scelefeed/util"
	"gorm.io/gorm"
)

type Scele struct {
	repo *repo.Scele
}

func NewSceleService(repo *repo.Scele) *Scele {
	return &Scele{
		repo: repo,
	}
}

func (svc *Scele) CreateUser(token, lineID string, sceleID int) (err error) {
	user := entity.UsersModel{
		Token:   token,
		LineId:  lineID,
		SceleID: sceleID,
	}
	err = svc.repo.CreateUser(user)
	return
}

func (svc *Scele) CreateNewCourse(token string, userID int, course entity.Course) (data entity.CoursesModel, err error) {
	var courseDetail []entity.ModulesResource
	data = entity.CoursesModel{}
	err = svc.repo.GetCourse(uint(course.Id), data)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	} else if err == gorm.ErrRecordNotFound {
		if err = svc.CreateTokenCourse(token, course.Id); err != nil {
			fmt.Println("aaaaa", err.Error())
			return
		}
		courseDetail, err = util.GetCourseDetail(token, course.Id)
		if err != nil {
			fmt.Println("bbbbb")
			return
		}
		err = svc.CreateCourse(token, userID, course, courseDetail)
	}
	return
}

func (svc *Scele) CreateCourse(token string, userID int, course entity.Course, courseDetail []entity.ModulesResource) (err error) {
	courseModel := entity.CoursesModel{
		CourseID:  uint(course.Id),
		ShortName: course.ShortName,
		LongName:  course.FullName,
		UserToken: token,
		Resource:  gin.H{"resource": courseDetail},
	}
	err = svc.repo.CreateCourse(courseModel)
	return
}

func (svc *Scele) CreateTokenCourse(token string, CourseID int) (err error) {
	tokenCourse := entity.TokenCourseModel{
		Token:    token,
		CourseID: uint(CourseID),
	}
	err = svc.repo.CreateTokenCourse(tokenCourse)
	return
}
