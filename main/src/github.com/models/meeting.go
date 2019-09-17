package models

type MeetingDetails struct {
	EmployeeId string `json:"employeeId"`
	SpokenTo []Details `json:"spokenTo"`
	Upcoming []Details `json:"upcomingMeeting"`
}

type Meeting struct{
	MeetingInfo []MeetingDetails `json:"meetingInfo"`
}

type Details struct{
	EmployeeId string `json:"employeeId"`
	MeetingDate string `json:"date"`
	MeetingTime string `json:"time"`
}
