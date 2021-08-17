package core

import (
	"fmt"
	"time"
)

func RunSchedule() {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fmt.Println("1 minute occured")
		}
	}
}	
