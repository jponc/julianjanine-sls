package schema

import "github.com/jponc/julianjanine/internal/types"

type HealthcheckResponse struct {
	Message string `json:"message"`
}

type GetGuestsResponse *[]types.Guest

type UpdateAttendanceRequest struct {
	Attendance types.Attendance `json:"attendance"`
}

type UpdateAttendanceResponse struct {
	Message string `json:"message"`
}
