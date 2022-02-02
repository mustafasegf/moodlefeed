package entity

type CourseResource struct {
	Id          int               `json:"id"`
	Name        string            `json:"name"`
	Visible     int               `json:"visible"`
	Uservisible bool              `json:"uservisible"`
	Modules     []ModulesResource `json:"modules"`
}
type ModulesResource struct {
	Id                  int                `json:"id"`
	Url                 string             `json:"url"`
	Name                string             `json:"name"`
	Instance            int                `json:"instance"`
	Description         string             `json:"description"`
	Visible             int                `json:"visible"`
	Uservisible         bool               `json:"uservisible"`
	Visibleoncoursepage int                `json:"visibleoncoursepage"`
	Modname             string             `json:"modname"`
	Contents            []ContentsResource `json:"contents"`
}

type ContentsResource struct {
	Type     string `json:"type"`
	FileName string `json:"filename"`
	FileUrl  string `json:"fileurl"`
}

type Course struct {
	ID        int    `json:"id"`
	ShortName string `json:"shortname"`
	FullName  string `json:"fullname"`
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
