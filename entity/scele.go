package entity

import (
	"database/sql"
	"time"
)

type Default struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
}

type Course struct {
	Id        int    `json:"id"`
	ShortName string `json:"shortname"`
	FullName  string `json:"fullname"`
}

type ContentsResource struct {
	Type     string `json:"type"`
	FileName string `json:"filename"`
	FileUrl  string `json:"fileurl"`
}
type ModulesResource struct {
	Id                  int                `json:"id"`
	Url                 string             `json:"url"`
	Name                string             `json:"name"`
	Description         string             `json:"description"`
	Visible             int                `json:"visible"`
	Uservisible         bool               `json:"uservisible"`
	Visibleoncoursepage int                `json:"visibleoncoursepage"`
	Modname             string             `json:"modname"`
	Contents            []ContentsResource `json:"contents"`
}

type CourseResource struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Visible     int               `json:"visible"`
	Uservisible bool              `json:"uservisible"`
	Modules     []ModulesResource `json:"modules"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Token struct {
	Token string `json:"token"`
}

type SceleUser struct {
	SceleID int `json:"userid"`
}

type UsersModel struct {
	Default
	Token   string
	LineId  string
	SceleID int
}

type CoursesModel struct {
	Default
	CourseID  uint
	ShortName string
	LongName  string
	UserToken string
	Resource  Resource
}

type Resource struct {
	Resource []CourseResource `json:"resource"`
}

// func (j *Resource) Scan(value interface{}) error {
// 	bytes, ok := value.([]byte)
// 	if !ok {
// 		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
// 	}

// 	result := Resource{}
// 	err := json.Unmarshal(bytes, &result)
// 	*j = Resource(result)
// 	return err
// }

// func (j *Resource) Value() (driver.Value, error) {
// 	if j == nil {
// 		return nil, nil
// 	}
// 	return &j, nil
// }

type MessageTypeModel struct {
	Default
	Name string
}

type TokenCourseModel struct {
	Default
	CourseID uint
	Token    string
}

type UserSubscribeModel struct {
	Default
	UserId   uint
	TypeId   uint
	CourseId uint
}
