package repository

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	awsDynamodb "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/jponc/julianjanine/internal/types"
	"github.com/jponc/julianjanine/pkg/dynamodb"
)

type guestItem struct {
	PK   string      `json:"PK"`
	SK   string      `json:"SK"`
	Data types.Guest `json:"Data"`
}

type Repository struct {
	dynamodbClient *dynamodb.Client
}

// NewClient instantiates a repository
func NewClient(dynamodbClient *dynamodb.Client) (*Repository, error) {
	r := &Repository{
		dynamodbClient: dynamodbClient,
	}

	return r, nil
}

func (r *Repository) GetGuests(ctx context.Context, invitationCode string) (*[]types.Guest, error) {
	items := []guestItem{}

	input := &awsDynamodb.QueryInput{
		KeyConditionExpression: aws.String("PK = :PK and begins_with(SK, :SK)"),
		ExpressionAttributeValues: map[string]*awsDynamodb.AttributeValue{
			":PK": {
				S: aws.String(fmt.Sprintf("Invite_%s", invitationCode)),
			},
			":SK": {
				S: aws.String("Guest_"),
			},
		},
		TableName: aws.String(r.dynamodbClient.GetTableName()),
	}

	output, err := r.dynamodbClient.Query(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to query guests: %v", err)
	}

	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &items)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal map: %v", err)
	}

	guests := []types.Guest{}
	for _, i := range items {
		guests = append(guests, i.Data)
	}

	return &guests, nil
}
