package core

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/httprequest"
	"github.com/mustafasegf/scelefeed/service"
)

type Schedule struct {
	svc *service.Scele
}

func NewSchedule(svc *service.Scele) *Schedule {
	return &Schedule{
		svc: svc,
	}
}

func (s *Schedule) RunSchedule() {
	s.GetCourse()
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("1 minute occured")
			s.GetCourse()
		}
	}
}

func (s *Schedule) GetCourse() {

	courses, _ := s.svc.GetAllCourse()

	for _, course := range courses {
		newCourseResource, _ := httprequest.GetCourseDetail(course.UserToken, int(course.CourseID))
		newCourse := entity.Resource{Resource: newCourseResource}
		if course.CourseID == 702 {
			fmt.Printf("%v\n", newCourse)
		}
		eq := cmp.Equal(newCourse, course.Resource)
		if !eq {
			err := s.svc.UpdateCourseResource(course.CourseID, newCourse)
			if err != nil {
				log.Printf("error updating course %s with course id %d. error code : %v", course.LongName, course.CourseID, err)
			}
		}
	}
}
