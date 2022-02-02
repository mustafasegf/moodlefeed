package entity

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gopkg.in/guregu/null.v4"
)

type UsersModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Token     string
	SceleID   int
}

type ClientModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	SceleID   int
	LineID    null.String
	DiscordID null.String
}

type CoursesModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	CourseID  int
	ShortName string
	LongName  string
	UserToken string
	Resource  Resource
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
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	Name      string
}

type TokenCourseModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	CourseID  int
	Token     string
}

type UserSubscribeModel struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	SceleID   int
	TypeID    int
	CourseID  int
}
