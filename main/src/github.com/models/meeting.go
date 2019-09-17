package models

type MeetingDetails struct {
	EmployeeId string `json:"employeeId"`
	SpokenTo []string `json:"spokenTo"`
	Upcoming []string `json:"spokenTo"`
}

type Meeting struct{
	MeetingInfo []MeetingDetails `json:"meetingInfo"`
}

type UpcomingDetails struct{
	EmployeeId string `json:"employeeId"`
	MeetingDate string `json:"date"`
	MeetingTime string `json:"time"`
}
