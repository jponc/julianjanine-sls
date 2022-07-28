package dynamodb

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-xray-sdk-go/xray"
)

type Client struct {
	dynamodbClient *dynamodb.DynamoDB
	tableName      string
}

// NewClient instantiates a DynamoDB Client
func NewClient(awsRegion, tableName string) (*Client, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion),
	})
	if err != nil {
		return nil, fmt.Errorf("cannot create aws session: %v", err)
	}

	dynamodbClient := dynamodb.New(sess)
	xray.AWS(dynamodbClient.Client)

	c := &Client{
		dynamodbClient: dynamodbClient,
		tableName:      tableName,
	}

	return c, nil
}

func (c *Client) Scan(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return c.dynamodbClient.ScanWithContext(ctx, input)
}

func (c *Client) PutItem(ctx context.Context, input *dynamodb.PutItemInput) (*dynamodb.PutItemOutput, error) {
	return c.dynamodbClient.PutItemWithContext(ctx, input)
}

func (c *Client) BatchWriteItem(ctx context.Context, input *dynamodb.BatchWriteItemInput) (*dynamodb.BatchWriteItemOutput, error) {
	return c.dynamodbClient.BatchWriteItemWithContext(ctx, input)
}

func (c *Client) Query(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return c.dynamodbClient.QueryWithContext(ctx, input)
}

func (c *Client) GetItem(ctx context.Context, input *dynamodb.GetItemInput) (*dynamodb.GetItemOutput, error) {
	return c.dynamodbClient.GetItemWithContext(ctx, input)
}

func (c *Client) GetTableName() string {
	return c.tableName
}
