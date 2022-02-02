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

func (svc *Service) CreateUser(token string, sceleID int) (err error) {
	user := entity.UsersModel{
		Token:   token,
		SceleID: sceleID,
	}
	err = svc.repo.CreateUser(user)
	return
}

func (svc *Service) CreateNewCourse(token string, userID int, course entity.Course) (data entity.CoursesModel, err error) {
	var courseDetail []entity.CourseResource

	httprequest := HttpRequest{}
	data, err = svc.repo.GetCourseByID(course.ID)
	if err != gorm.ErrRecordNotFound && err != nil {
		return
	} else if err == gorm.ErrRecordNotFound {
		if courseDetail, err = httprequest.GetCourseDetail(token, course.ID); err != nil {
			return
		}
		if err = svc.CreateCourse(token, userID, course, courseDetail); err != nil {
			return
		}
		if err = svc.CreateTokenCourse(token, course.ID); err != nil {
			return
		}
		err = svc.DefaultSubscribe(userID, course.ID)
	}
	return
}

func (svc *Service) CreateCourse(token string, userID int, course entity.Course, courseDetail []entity.CourseResource) (err error) {
	courseModel := entity.CoursesModel{
		CourseID:  course.ID,
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
		CourseID: CourseID,
	}
	err = svc.repo.CreateTokenCourse(tokenCourse)
	return
}

func (svc *Service) DeleteUserCourse(sceleID int, course entity.Course) (err error) {
	// userSubscribeModel := entity.UserSubscribeModel{
	// 	SceleID: sceleID,
	// }
	// err = svc.repo.DeleteCourse(userSubscribeModel)
	return
}

func (svc *Service) DefaultSubscribe(sceleID, CourseID int) (err error) {
	for i := 1; i < 11; i++ {
		userSubscribe := entity.UserSubscribeModel{
			SceleID:   sceleID,
			CourseID: CourseID,
			TypeID:   i,
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

func (svc *Service) UpdateCourseResource(courseID int, resource entity.Resource) (err error) {
	model := entity.CoursesModel{
		CourseID: courseID,
		Resource: resource,
	}
	err = svc.repo.UpdateCourseResource(model)
	return
}

func (svc *Service) GetLineIDFromCourse(courseID int) (user []entity.UsersModel, err error) {
	user, err = svc.repo.GetLineIDFromCourse(courseID)
	return
}

func (svc *Service) GetCoursesNameByToken(token string) (courses []entity.Course, err error) {
	courses, err = svc.repo.GetCoursesNameByToken(token)
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

	res = "removed:\n"
	for _, course := range oldCourses {
		oldCoursesSet[course.ShortName] = course
		if _, exist := newCourseSet[course.ShortName]; !exist {
			deletedCourse = append(deletedCourse, course.ShortName)
			res += "- " + course.ShortName + "\n"
		}
	}

	if len(deletedCourse) == 0 {
		res += "none\n"
	}

	res += "added:\n"
	for _, course := range newCourses {
		if _, exist := oldCoursesSet[course.ShortName]; !exist {
			createdCourse = append(createdCourse, course.ShortName)
			res += "- " + course.ShortName + "\n"
		}
	}

	if len(createdCourse) == 0 {
		res += "none\n"
	}

	for _, course := range createdCourse {
		_, err = svc.CreateNewCourse(token, sceleID, newCourseSet[course])
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
	}

	//TODO: added delete course
	for _, course := range deletedCourse {
		err = svc.DeleteUserCourse(sceleID, oldCoursesSet[course])
		if err != nil {
			fmt.Printf("err:%v\n", err)
		}
	}

	return
}
