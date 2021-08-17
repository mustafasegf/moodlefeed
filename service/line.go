package service

import (
	"fmt"
	"os"
)

type Line struct{}

func NewLineService() *Line {
	return &Line{}
}

func (svc *Line) GetLoginUrl(lineID string) (url string) {
	url = fmt.Sprintf("%s?id=%s", os.Getenv("SERVER_HOST"), lineID)
	return
}
