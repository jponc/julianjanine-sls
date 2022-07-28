package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jponc/julianjanine/internal/api"
	"github.com/jponc/julianjanine/internal/repository"
	"github.com/jponc/julianjanine/pkg/dynamodb"
	log "github.com/sirupsen/logrus"
)

func main() {
	config, err := NewConfig()
	if err != nil {
		log.Fatalf("cannot initialise config %v", err)
	}

	dynamodbClient, err := dynamodb.NewClient(config.AWSRegion, config.DBTableName)
	if err != nil {
		log.Fatalf("cannot initialise dynamodb client %v", err)
	}

	repository, err := repository.NewClient(dynamodbClient)
	if err != nil {
		log.Fatalf("cannot initialise repository %v", err)
	}

	service := api.NewService(repository)
	lambda.Start(service.GetGuests)
}
