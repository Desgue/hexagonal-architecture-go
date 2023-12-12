package vehiclerepo

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

func createTable(dynamoRepository *dynamoRepository) error {
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
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String(dynamoRepository.tableName),
	}
	_, err := dynamoRepository.client.CreateTable(input)
	if err != nil {
		return err

	}
	return nil
}

func deleteTable(dynamoRepository *dynamoRepository) error {
	dynamoRepository.client.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: &dynamoRepository.tableName,
	})
	return nil
}

func populateDb(dynamoRepository dynamoRepository) ([]domain.Vehicle, error) {
	raw, err := os.ReadFile("vehicles.json")
	if err != nil {
		return []domain.Vehicle{}, err
	}

	var vehicles []domain.Vehicle
	err = json.Unmarshal(raw, &vehicles)
	if err != nil {
		log.Println("Error Unmarshaling json: ", err.Error())
		return []domain.Vehicle{}, err

	}

	for _, member := range vehicles {
		av, err := dynamodbattribute.MarshalMap(member)
		if err != nil {
			return []domain.Vehicle{}, err
		}
		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(dynamoRepository.tableName),
		}
		_, err = dynamoRepository.client.PutItem(input)
		if err != nil {
			return []domain.Vehicle{}, err
		}
	}
	return vehicles, nil
}

type testContainer struct {
	*localstack.LocalStackContainer
	URI string
}

func prepareContainer(ctx context.Context) (*testContainer, error) {
	container, err := localstack.RunContainer(ctx, testcontainers.WithImage("localstack/localstack:latest"))
	if err != nil {
		return nil, err
	}
	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		return nil, err
	}
	host, err := provider.DaemonHost(ctx)
	if err != nil {
		return nil, err
	}
	mappedPort, err := container.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		return nil, err
	}
	URI := fmt.Sprintf("http://%s:%d", host, mappedPort.Int())
	return &testContainer{LocalStackContainer: container, URI: URI}, nil
}
