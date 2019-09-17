package models

type Employee struct {
	Name     string `json:"name"`
	Team     string `json:"team"`
	About    string `json:"aboutMe"`
	Location string `json:"location"`
	Email    string `json:"email"`
}
