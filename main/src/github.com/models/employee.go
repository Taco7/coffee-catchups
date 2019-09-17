package models

type EmployeeDetails struct {
	EmployeeId	string `json:"employeeId"`
	Name     string `json:"name"`
	Team     string `json:"team"`
	About    string `json:"aboutMe"`
	Location string `json:"location"`
	Email    string `json:"email"`
}

type Employees struct{
	Employee []EmployeeDetails `json:"employee"`
}