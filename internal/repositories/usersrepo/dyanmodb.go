package usersrepo

import (
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type dynamoRepository struct {
	client    *dynamodb.DynamoDB
	tableName string
}

func configDB() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region:   aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:8080"),
		},
		Profile: "default",
	}))
	client := dynamodb.New(sess)
	return client

}
func NewDynamoRepository(tableName string) *dynamoRepository {
	return &dynamoRepository{
		client:    configDB(),
		tableName: tableName,
	}
}

// By defining these two methods we implement the UserRepository interface from ports
func (d *dynamoRepository) Save() error {
	return nil
}
func (d *dynamoRepository) FindAll() ([]domain.User, error) {
	return nil, nil
}
