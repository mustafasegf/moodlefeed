package controller

import (
	"fmt"
	"net/http"

	"github.com/mustafasegf/scelefeed/service"
)

type Scele struct {
	svc *service.Scele
}

func NewSceleController(svc *service.Scele) *Scele {
	return &Scele{
		svc: svc,
	}
}

func (ctrl *Scele) Home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}
