package api

import (
	"context"

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
