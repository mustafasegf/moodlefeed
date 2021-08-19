package service

import (
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/httprequest"
	"github.com/mustafasegf/scelefeed/repo"
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
	var courseDetail []entity.CourseResource
	data = entity.CoursesModel{}
	err = svc.repo.GetCourse(uint(course.Id), data)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	} else if err == gorm.ErrRecordNotFound {
		if courseDetail, err = httprequest.GetCourseDetail(token, course.Id); err != nil {
			return
		}
		if err = svc.CreateCourse(token, userID, course, courseDetail); err != nil {
			return
		}
		if err = svc.CreateTokenCourse(token, course.Id); err != nil {
			return
		}
		err = svc.DefaultSubscribe(userID, course.Id)
	}
	return
}

func (svc *Scele) CreateCourse(token string, userID int, course entity.Course, courseDetail []entity.CourseResource) (err error) {
	courseModel := entity.CoursesModel{
		CourseID:  uint(course.Id),
		ShortName: course.ShortName,
		LongName:  course.FullName,
		UserToken: token,
		Resource:  entity.Resource{Resource: courseDetail},
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

func (svc *Scele) DefaultSubscribe(userID, CourseID int) (err error) {
	var i uint
	for i = 1; i < 11; i++ {
		userSubscribe := entity.UserSubscribeModel{
			UserId:   uint(userID),
			CourseId: uint(CourseID),
			TypeId:   i,
		}
		err = svc.repo.CreateUserSubscribe(userSubscribe)
		if err != nil {
			return
		}
	}

	return
}

func (svc *Scele) GetAllCourse() (courses []entity.CoursesModel, err error) {
	courses = []entity.CoursesModel{}
	svc.repo.GetAllCourse(&courses)

	return
}
