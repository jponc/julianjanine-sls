package api

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/jponc/julianjanine/internal/schema"
	"github.com/jponc/julianjanine/pkg/lambdaresponses"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Healthcheck(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return lambdaresponses.Respond200(schema.HealthcheckResponse{Message: "OK"})
}
