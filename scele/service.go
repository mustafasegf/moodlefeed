package scele

import (
	"fmt"

	"github.com/mustafasegf/scelefeed/entity"
	"gorm.io/gorm"
)

type Service struct {
	repo *Repo
}

var httprequest *HttpRequest

// i know this is kinda ugly. will refactor later
func init() {
	httprequest = &HttpRequest{}
}

func NewService(repo *Repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (svc *Service) CreateUser(token, lineID string, sceleID int) (err error) {
	user := entity.UsersModel{
		Token:   token,
		LineId:  lineID,
		SceleID: sceleID,
	}
	err = svc.repo.CreateUser(user)
	return
}

func (svc *Service) CreateNewCourse(token string, userID int, course entity.Course) (data entity.CoursesModel, err error) {
	var courseDetail []entity.CourseResource
	data = entity.CoursesModel{}
	httprequest := HttpRequest{}
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

func (svc *Service) CreateCourse(token string, userID int, course entity.Course, courseDetail []entity.CourseResource) (err error) {
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

func (svc *Service) CreateTokenCourse(token string, CourseID int) (err error) {
	tokenCourse := entity.TokenCourseModel{
		Token:    token,
		CourseID: uint(CourseID),
	}
	err = svc.repo.CreateTokenCourse(tokenCourse)
	return
}

func (svc *Service) DefaultSubscribe(userID, CourseID int) (err error) {
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

func (svc *Service) GetAllCourse() (courses []entity.CoursesModel, err error) {
	courses = []entity.CoursesModel{}
	svc.repo.GetAllCourse(&courses)

	return
}

func (svc *Service) UpdateCourseResource(courseId uint, resource entity.Resource) (err error) {
	err = svc.repo.UpdateCourseResource(courseId, resource)
	return
}

func (svc *Service) GetIdLineFromCourse(courseId uint) (user []entity.UsersModel, err error) {
	user = []entity.UsersModel{}
	err = svc.repo.GetIdLineFromCourse(courseId, &user)
	return
}

func (svc *Service) GetCoursesNameByToken(token string) (courses []entity.Course, err error) {
	courses = []entity.Course{}
	err = svc.repo.GetCoursesNameByToken(token, &courses)
	return
}

func (svc *Service) UpdateUserCourse(token string, sceleID int) (res string, err error) {

	// get all course
	newCourses, err := httprequest.GetCourses(token, sceleID)
	if err != nil {
		return
	}

	newCourseSet := make(map[string]entity.Course)
	for _, course := range newCourses {
		newCourseSet[course.ShortName] = course
	}

	oldCourses, err := svc.GetCoursesNameByToken(token)
	if err != nil {
		return
	}

	oldCoursesSet := make(map[string]entity.Course)

	createdCourse := make([]string, 0)
	deletedCourse := make([]string, 0)

	for _, course := range oldCourses {
		oldCoursesSet[course.ShortName] = course
		if _, exist := newCourseSet[course.ShortName]; !exist {
			deletedCourse = append(deletedCourse, course.ShortName)
		}
	}

	for _, course := range newCourses {
		if _, exist := oldCoursesSet[course.ShortName]; !exist {
			createdCourse = append(createdCourse, course.ShortName)
		}
	}

	for _, course := range createdCourse {
		_, err = svc.CreateNewCourse(token, sceleID, newCourseSet[course])
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
	}
	//TODO: added delete course
	return
}
