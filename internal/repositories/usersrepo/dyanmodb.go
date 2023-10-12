package usersrepo

import "github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"

type dynamoRepository struct {
}

func NewDynamoRepository() *dynamoRepository {
	return &dynamoRepository{}
}

// By defining these two methods we implement the UserRepository interface from ports
func (d *dynamoRepository) Save() error {
	return nil
}
func (d *dynamoRepository) FindAll() ([]domain.User, error) {
	return nil, nil
}
