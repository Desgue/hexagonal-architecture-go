package usersrepo

import (
	"errors"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
)

type FakeRepository struct {
	memoryStorage map[string]domain.User
}

func NewFakeRepository(store map[string]domain.User) *FakeRepository {
	return &FakeRepository{memoryStorage: store}
}

func (this *FakeRepository) Insert(user domain.User) (domain.User, error) {
	this.memoryStorage[user.Id] = user
	return user, nil
}

func (this *FakeRepository) FindAll() ([]domain.User, error) {
	allUsers := []domain.User{}

	if len(this.memoryStorage) == 0 {
		return nil, errors.New("Error retrieving users: No user found")
	}
	for _, value := range this.memoryStorage {
		allUsers = append(allUsers, value)
	}
	return allUsers, nil
}

func (this *FakeRepository) FindById(id string) (domain.User, error) {
	user, ok := this.memoryStorage[id]
	if !ok {
		return domain.User{}, errors.New("User not found")
	}
	return user, nil
}
