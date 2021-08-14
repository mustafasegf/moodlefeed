package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mustafasegf/scelefeed/entity"
)

func LoginScele(username, password string) (token entity.Token, err error) {
	url := fmt.Sprintf("https://scele.cs.ui.ac.id/login/token.php?service=moodle_mobile_app&username=%s&password=%s", username, password)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	token = entity.Token{}
	err = json.NewDecoder(resp.Body).Decode(&token)
	return
}
