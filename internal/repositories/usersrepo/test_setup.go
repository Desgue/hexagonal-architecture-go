package usersrepo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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
