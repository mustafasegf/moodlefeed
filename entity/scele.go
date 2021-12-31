package entity

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

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
	gorm.Model
	Token   string `gorm:"column:token"`
	LineId  string `gorm:"column:line_id"`
	SceleID int    `gorm:"column:scele_id"`
}

type CoursesModel struct {
	gorm.Model
	CourseID  uint     `gorm:"column:course_id"`
	ShortName string   `gorm:"column:short_name"`
	LongName  string   `gorm:"column:long_name"`
	UserToken string   `gorm:"column:user_token"`
	Resource  Resource `gorm:"column:resource;type:json" json:"resouce"`
}

type Resource struct {
	Resource []CourseResource `json:"resource"`
}

func (j *Resource) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}

	result := Resource{}
	err := json.Unmarshal(bytes, &result)
	*j = Resource(result)
	return err
}

func (j *Resource) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return &j, nil
}

type MessageTypeModel struct {
	gorm.Model
	Name string `gorm:"column:name"`
}

type TokenCourseModel struct {
	gorm.Model
	CourseID uint   `gorm:"column:course_id"`
	Token    string `gorm:"column:token"`
}

type UserSubscribeModel struct {
	gorm.Model
	UserId   uint `gorm:"column:user_id"`
	TypeId   uint `gorm:"column:type_id"`
	CourseId uint `gorm:"column:course_id"`
}
