package client

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type Line struct {
	bot *linebot.Client
}

func NewLineController(bot *linebot.Client) *Line {
	return &Line{
		bot: bot,
	}
}

func (ctrl *Line) GetLoginUrl(lineID string) (url string) {
	url = fmt.Sprintf("%s?id=%s", os.Getenv("SERVER_HOST"), lineID)
	return
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
					res := ctrl.GetLoginUrl(event.Source.UserID)
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
