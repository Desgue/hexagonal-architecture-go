package usersrepo

import (
	"errors"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type dynamoRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func NewDynamoRepository(endpoint, tableName string) *dynamoRepository {
	return &dynamoRepository{
		client:    configClientDB(endpoint),
		tableName: tableName,
	}
}

// By defining these two methods we implement the UserRepository interface from ports
func (this *dynamoRepository) Insert(user domain.User) (domain.User, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		return domain.User{}, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(this.tableName),
	}
	_, err = this.client.PutItem(input)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (this *dynamoRepository) FindById(id string) (domain.User, error) {
	result, err := this.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(this.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return domain.User{}, err
	}
	if result.Item == nil {
		return domain.User{}, errors.New("No user found with given ID")
	}
	foundUser := domain.User{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &foundUser)
	if err != nil {
		return domain.User{}, err
	}
	return foundUser, nil
}

// Private functions

func (this *dynamoRepository) createTable() error {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
		},
		TableName:   aws.String(this.tableName),
		BillingMode: aws.String("PAY_PER_REQUEST"),
	}
	_, err := this.client.CreateTable(input)
	if err != nil {
		return err

	}
	return nil
}

func configClientDB(endpoint string) *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String(endpoint),
		},
		Profile: "default",
	}))
	client := dynamodb.New(sess)
	return client

}
