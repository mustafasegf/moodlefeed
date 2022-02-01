package line

import (
	"fmt"
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/scelefeed/entity"
	"github.com/mustafasegf/scelefeed/scele"
)

type Controller struct {
	Bot      *linebot.Client
	Svc      *Service
	SceleSvc *scele.Service
}

func NewController(bot *linebot.Client, svc *Service, sceleSvc *scele.Service) *Controller {
	return &Controller{
		Bot:      bot,
		Svc:      svc,
		SceleSvc: sceleSvc,
	}
}

func (ctrl *Controller) LineCallback(w http.ResponseWriter, req *http.Request) {
	events, err := ctrl.Bot.ParseRequest(req)
	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	for _, event := range events {
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				switch message.Text {
				case "/login":
					res := ctrl.Svc.GetLoginUrl(event.Source.UserID)
					if _, err = ctrl.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				case "/update":
					// get scele credentials
					user, err := ctrl.Svc.GetUserFromLineID(event.Source.UserID)
					if err != nil {
						log.Println(err)
						continue
					}

					// get all course
					/* courses, err := httprequest.GetCourses(user.Token, user.SceleID)
					if err != nil {
						log.Println(err)
						continue
					} */

					// save to db
					// TODO: if new course is not on db, delete on db
					ctrl.SceleSvc.UpdateUserCourse(user.Token, user.SceleID)
					newCourse := []entity.Course{}
					/* for _, course := range courses {
						_, err = ctrl.scelesvc.CreateNewCourse(user.Token, user.SceleID, course)
						if err != nil {
							if !strings.Contains(err.Error(), "unique") {
								log.Println("err")
								continue
							}
						} else {
							newCourse = append(newCourse, course)
						}
					} */

					if _, err = ctrl.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(fmt.Sprintf("%v", newCourse))).Do(); err != nil {
						log.Print(err)
					}
				case "/courses":
					res := ""

					user, err := ctrl.Svc.GetUserFromLineID(event.Source.UserID)
					if err != nil {
						log.Println(err)
						continue
					}

					courses, err := ctrl.SceleSvc.GetCoursesNameByToken(user.Token)
					if err != nil {
						res = err.Error()
					} else {
						for _, course := range courses {
							res += course.ShortName + "\n"
						}
					}

					if _, err = ctrl.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				case "/help":
					res := `/login
/update
/course"`
					if _, err = ctrl.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}

				default:
					res := "use /help"
					if _, err = ctrl.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
