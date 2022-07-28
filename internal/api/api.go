package api

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jponc/julianjanine/internal/schema"
	"github.com/jponc/julianjanine/internal/types"
	"github.com/jponc/julianjanine/pkg/lambdaresponses"
	log "github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
}

type Repository interface {
	GetGuests(ctx context.Context, invitationCode string) (*[]types.Guest, error)
	UpdateAttendance(ctx context.Context, invitationCode string, guestId string, attendance types.Attendance) error
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) Healthcheck(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return lambdaresponses.Respond200(schema.HealthcheckResponse{Message: "OK"})
}

func (s *Service) GetGuests(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if s.repository == nil {
		log.Errorf("repository is missing")
		return lambdaresponses.Respond500()
	}

	invitationCode := request.PathParameters["invitationCode"]
	if invitationCode == "" {
		log.Errorf("invitationCode missing")
		return lambdaresponses.Respond500()
	}

	guests, err := s.repository.GetGuests(ctx, invitationCode)
	if err != nil {
		log.Errorf("error when fetching guests for code %s: %w", invitationCode, err)
		return lambdaresponses.Respond500()
	}

	resp := schema.GetGuestsResponse(guests)
	return lambdaresponses.Respond200(resp)
}

func (s *Service) UpdateAttendance(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if s.repository == nil {
		log.Errorf("repository is missing")
		return lambdaresponses.Respond500()
	}

	invitationCode := request.PathParameters["invitationCode"]
	if invitationCode == "" {
		log.Errorf("invitationCode missing")
		return lambdaresponses.Respond500()
	}

	guestId := request.PathParameters["guestId"]
	if guestId == "" {
		log.Errorf("guestId missing")
		return lambdaresponses.Respond500()
	}

	var req schema.UpdateAttendanceRequest
	err := json.Unmarshal([]byte(request.Body), &req)

	if req.Attendance != types.AttendanceNo && req.Attendance != types.AttendanceYes && req.Attendance != types.AttendanceTentative {
		log.Errorf("attendance passed is not supported: %s", req.Attendance)
		return lambdaresponses.Respond500()
	}

	err = s.repository.UpdateAttendance(ctx, invitationCode, guestId, req.Attendance)
	if err != nil {
		log.Errorf("error updating attendance (%s) for guest (%s): %w", req.Attendance, guestId, err)
		return lambdaresponses.Respond500()
	}

	return lambdaresponses.Respond200(schema.UpdateAttendanceResponse{Message: "OK"})
}
