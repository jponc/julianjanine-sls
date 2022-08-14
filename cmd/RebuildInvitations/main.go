package main

import (
	_ "embed"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jponc/julianjanine/internal/invitationscsv"
	"github.com/jponc/julianjanine/internal/repository"
	"github.com/jponc/julianjanine/internal/tools"
	"github.com/jponc/julianjanine/pkg/dynamodb"
	log "github.com/sirupsen/logrus"
)

//go:embed data.csv
var csvData []byte

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

	invitationsCSV := invitationscsv.NewClient(csvData)

	tool := tools.NewClient(repository, invitationsCSV)
	lambda.Start(tool.RebuildInvitations)
}
