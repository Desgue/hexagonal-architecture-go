package vehiclerepo

import (
	"errors"
	"fmt"
	"log"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/pkg/apperr"
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

func (db *dynamoRepository) Insert(vehicle domain.Vehicle) (domain.Vehicle, error) {
	entityParsed, err := dynamodbattribute.MarshalMap(vehicle)
	if err != nil {
		return domain.Vehicle{}, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(db.tableName),
	}
	_, err = db.client.PutItem(input)
	if err != nil {
		return domain.Vehicle{}, err
	}
	return vehicle, nil
}
func (db *dynamoRepository) FindAll() ([]domain.Vehicle, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(db.tableName),
	}
	result, err := db.client.Scan(input)

	if err != nil {
		return []domain.Vehicle{}, err
	}

	if len(result.Items) == 0 {
		return []domain.Vehicle{}, apperr.NotFound
	}

	var users []domain.Vehicle

	for _, i := range result.Items {
		vehicle := domain.Vehicle{}
		err = dynamodbattribute.UnmarshalMap(i, &vehicle)
		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err)
			return nil, err
		}
		users = append(users, vehicle)
	}
	return users, nil

}
func (db *dynamoRepository) FindById(id string) (domain.Vehicle, error) {
	result, err := db.client.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(db.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		return domain.Vehicle{}, err
	}
	if result.Item == nil {
		return domain.Vehicle{}, errors.New("No vehicle found with given ID")
	}
	foundUser := domain.Vehicle{}
	err = dynamodbattribute.UnmarshalMap(result.Item, &foundUser)
	if err != nil {
		return domain.Vehicle{}, err
	}
	return foundUser, nil
}

func (db *dynamoRepository) Delete(id string) error {

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(id),
			},
		},
		TableName: aws.String(db.tableName),
	}

	_, err := db.client.DeleteItem(input)
	if err != nil {
		log.Printf("Error calling DeleteItem: %s", err)
		return err
	}
	log.Println("Successfully deleted item from database")
	return nil
}

// Private functions
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
