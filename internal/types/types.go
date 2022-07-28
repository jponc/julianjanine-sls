package types

type Invitation struct {
	Code string `json:"code"`
}

type Attendance string

const (
	AttendancePending   Attendance = "Pending"
	AttendanceTentative Attendance = "Tentative"
	AttendanceYes       Attendance = "Yes"
	AttendanceNo        Attendance = "No"
)

type Guest struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Attendance Attendance `json:"attendance"`
}
