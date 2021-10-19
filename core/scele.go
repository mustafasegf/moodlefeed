package core

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/httprequest"
	"github.com/mustafasegf/scelefeed/service"
	"github.com/mustafasegf/scelefeed/util"
)

type Schedule struct {
	svc *service.Scele
	bot *linebot.Client
}

func NewSchedule(svc *service.Scele, bot *linebot.Client) *Schedule {
	return &Schedule{
		svc: svc,
		bot: bot,
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
			// s.GetCourse()
		}
	}
}

func (s *Schedule) GetCourse() {

	courses, _ := s.svc.GetAllCourse()

	for _, course := range courses {
		// if course.CourseID != 3197 {
		// 	continue
		// }
		newCourseResource, _ := httprequest.GetCourseDetail(course.UserToken, int(course.CourseID))
		newCourse := entity.Resource{Resource: newCourseResource}
		var r util.DiffReporter
		eq := cmp.Equal(newCourse, course.Resource, cmp.Reporter(&r))
		if !eq {
			diff := r.GetDiff()
			msg := make([]string, 0, len(diff))
			for _, index := range diff {
				res := newCourse.Resource[index.Resource].Modules[index.Modules]
				tmp := fmt.Sprintf("%s\n%s\n%s\n\n", res.Name, res.Description, res.Url)
				fmt.Println(tmp)
				msg = append(msg, tmp)
			}

			user, err := s.svc.GetIdLineFromCourse(course.CourseID)
			if err != nil {
				continue
			}
			for _, user := range user {
				s.Message(user.LineId, strings.Join(msg, "\n"))
			}
			err = s.svc.UpdateCourseResource(course.CourseID, newCourse)
			if err != nil {
				log.Printf("error updating course %s with course id %d. error code : %v", course.LongName, course.CourseID, err)
				return
			}

		}
	}
}

func (s *Schedule) Message(idLine, message string) {
	res := s.bot.PushMessage(idLine, linebot.NewTextMessage(message))
	_, err := res.Do()
	fmt.Printf(">>> err: %v\n", err)
}
