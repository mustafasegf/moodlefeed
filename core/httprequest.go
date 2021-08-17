package core

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func RequestScele(token string, wsfunction string, args map[string]interface{}, model interface{}) (err error) {
	url := fmt.Sprintf("https://scele.cs.ui.ac.id/webservice/rest/server.php?moodlewsrestformat=json&wstoken=%s&wsfunction=%s", token, wsfunction)
	for k, v := range args {
		url = fmt.Sprintf("%s&%s=%v", url, k, v)
	}
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&model)
	return
}

func GetSceleId(token string) (sceleUser entity.SceleUser, err error) {
	sceleUser = entity.SceleUser{}
	err = RequestScele(token, "core_webservice_get_site_info", nil, &sceleUser)
	return
}

func GetCourses(token string, userid int) (courses []entity.Course, err error) {
	courses = make([]entity.Course, 0)
	err = RequestScele(token, "core_enrol_get_users_courses", gin.H{"userid": userid}, &courses)
	return
}

func GetCourseDetail(token string, courseID int) (resource []entity.CourseResource, err error) {
	resource = make([]entity.CourseResource, 0)
	err = RequestScele(token, "core_course_get_contents", gin.H{"courseid": courseID}, &resource)
	return
}
