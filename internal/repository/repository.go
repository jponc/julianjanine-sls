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

func (r *Repository) UpdateAttendance(ctx context.Context, invitationCode string, guestId string, attendance types.Attendance) error {
	item := guestItem{}

	// Gets the item
	input := &awsDynamodb.GetItemInput{
		Key: map[string]*awsDynamodb.AttributeValue{
			"PK": {
				S: aws.String(fmt.Sprintf("Invite_%s", invitationCode)),
			},
			"SK": {
				S: aws.String(fmt.Sprintf("Guest_%s", guestId)),
			},
		},
		TableName: aws.String(r.dynamodbClient.GetTableName()),
	}

	output, err := r.dynamodbClient.GetItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to get guest: %v", err)
	}

	if output.Item == nil {
		return fmt.Errorf("guest not found: invite (%s), guestId (%s)", invitationCode, guestId)
	}

	err = dynamodbattribute.UnmarshalMap(output.Item, &item)
	if err != nil {
		return fmt.Errorf("failed to unmarshal map: %v", err)
	}

	// Set new attendance
	guest := item.Data
	guest.Attendance = attendance

	// Update the guest record
	err = r.SaveGuest(ctx, invitationCode, guestId, guest)
	if err != nil {
		return fmt.Errorf("failed to save guest: %v", err)
	}

	return nil
}

func (r *Repository) SaveGuest(ctx context.Context, invitationCode string, guestId string, guest types.Guest) error {
	item := guestItem{
		PK:   fmt.Sprintf("Invite_%s", invitationCode),
		SK:   fmt.Sprintf("Guest_%s", guestId),
		Data: guest,
	}

	itemMap, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return fmt.Errorf("failed to ddb marshal result item record, %v", err)
	}

	input := &awsDynamodb.PutItemInput{
		Item:      itemMap,
		TableName: aws.String(r.dynamodbClient.GetTableName()),
	}

	_, err = r.dynamodbClient.PutItem(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to put Guest: %v", err)
	}

	return nil
}
