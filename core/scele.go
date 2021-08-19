package core

import (
	"encoding/json"
	"fmt"
	"time"

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
		}
	}
}

func (s *Schedule) GetCourse() {
	courses, _ := s.svc.GetAllCourse()

	b, _ := json.MarshalIndent(courses[0].LongName, "", "  ")
	fmt.Printf("%s\n", b)
}

// add check different
