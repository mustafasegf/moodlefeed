package core

import (
	"fmt"
	"html"
	"log"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/k3a/html2text"
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
	minute := 15
	ticker := time.NewTicker(time.Minute * time.Duration(minute))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println(minute, " minute occured")
			s.GetCourse()
		}
	}
}

func (s *Schedule) GetCourse() {

	courses, _ := s.svc.GetAllCourse()

	for _, course := range courses {
		go func(c entity.CoursesModel) {
			newCourseResource, _ := httprequest.GetCourseDetail(c.UserToken, int(c.CourseID))
			newCourse := entity.Resource{Resource: newCourseResource}
			var r util.DiffReporter
			eq := cmp.Equal(newCourse, c.Resource, cmp.Reporter(&r))
			if !eq {
				diff := r.GetDiff()
				if len(diff) == 0 {
					return
				}
				msg := make([]string, 0, len(diff)+1)
				msg = append(msg, fmt.Sprintf("%s\n\n", html.UnescapeString(c.LongName)))
				for _, index := range diff {
					res := newCourse.Resource[index.Resource].Modules[index.Modules]
					var tmp string
					// todo: add more modname custom string
					if res.Modname == "label" {
						tmp = fmt.Sprintf("%s\n\n", html2text.HTML2Text(res.Description))
					} else {
						tmp = fmt.Sprintf("%s\n%s\n%s\n\n", html.UnescapeString(res.Name), html2text.HTML2Text(res.Description), res.Url)
					}
					msg = append(msg, tmp)
				}
				fmt.Println(msg)
				user, err := s.svc.GetIdLineFromCourse(c.CourseID)
				if err != nil {
					return
				}
				for _, user := range user {
					s.Message(user.LineId, strings.Join(msg, "---\n"))
				}
				err = s.svc.UpdateCourseResource(c.CourseID, newCourse)
				if err != nil {
					log.Printf("error updating course %s with course id %d. error code : %v", c.LongName, c.CourseID, err)
					return
				}

			}
		}(course)
	}
}

func (s *Schedule) Message(idLine, message string) {
	res := s.bot.PushMessage(idLine, linebot.NewTextMessage(message))
	_, err := res.Do()
	if err != nil {
		log.Printf("cant push message: %v", err)
	}
}
