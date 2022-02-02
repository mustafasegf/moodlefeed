package line

import (
	"fmt"
	"os"

	"github.com/mustafasegf/scelefeed/entity"
)

type Service struct {
	Repo *Repo
}

func NewService(repo *Repo) *Service {
	return &Service{
		Repo: repo,
	}
}
func (svc *Service) GetLoginUrl(lineID string) (url string) {
	url = fmt.Sprintf("%s?id=%s", os.Getenv("SERVER_HOST"), lineID)
	return
}

func (svc *Service) GetUserFromLineID(lineID string) (user entity.UsersModel, err error) {
	// user = entity.UsersModel{
	// 	LineId: lineID,
	// }
	// err = svc.Repo.GetUserFromLineId(&user)
	return
}
