package scele

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mustafasegf/scelefeed/entity"
)

type HttpRequest struct {}

func (r* HttpRequest) LoginScele(username, password string) (token string, err error) {
	url := fmt.Sprintf("https://scele.cs.ui.ac.id/login/token.php?service=moodle_mobile_app&username=%s&password=%s", username, password)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	tokenEntity := entity.Token{}
	err = json.NewDecoder(resp.Body).Decode(&token)
	token = tokenEntity.Token
	return
}

func (r* HttpRequest) RequestScele(token string, wsfunction string, args map[string]interface{}, model interface{}) (err error) {
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

func (r* HttpRequest) GetSceleId(token string) (sceleUser entity.SceleUser, err error) {
	sceleUser = entity.SceleUser{}
	err = r.RequestScele(token, "core_webservice_get_site_info", nil, &sceleUser)
	return
}

func (r* HttpRequest) GetCourses(token string, sceleID int) (courses []entity.Course, err error) {
	courses = make([]entity.Course, 0)
	err = r.RequestScele(token, "core_enrol_get_users_courses", gin.H{"userid": sceleID}, &courses)
	return
}

func (r* HttpRequest) GetCourseDetail(token string, courseID int) (sanitizedResources []entity.CourseResource, err error) {
	resource := make([]entity.CourseResource, 0)
	err = r.RequestScele(token, "core_course_get_contents", gin.H{"courseid": courseID}, &resource)

	sanitizedResources = make([]entity.CourseResource, 0, len(resource))
	for _, r := range resource {
		if r.Uservisible && r.Visible == 1 {
			sanitizedModulesResource := make([]entity.ModulesResource, 0, len(r.Modules))

			for _, m := range r.Modules {
				if m.Uservisible && m.Visible == 1 && m.Visibleoncoursepage == 1 {
					sanitizedModulesResource = append(sanitizedModulesResource, m)
				}
			}
			r.Modules = sanitizedModulesResource
			sanitizedResources = append(sanitizedResources, r)
		}
	}
	return
}
