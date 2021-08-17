package controller

import (
	"log"
	"net/http"

	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/mustafasegf/scelefeed/service"
)

type Line struct {
	svc *service.Line
	bot *linebot.Client
}

func NewLineController(svc *service.Line, bot *linebot.Client) *Line {
	return &Line{
		svc: svc,
		bot: bot,
	}
}

func (ctrl *Line) LineCallback(w http.ResponseWriter, req *http.Request) {
	events, err := ctrl.bot.ParseRequest(req)
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
					res := ctrl.svc.GetLoginUrl(event.Source.UserID)
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				default:
					res := message.Text
					if _, err = ctrl.bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(res)).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	}
}
