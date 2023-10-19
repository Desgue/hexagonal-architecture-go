package userservice

import (
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/ports"
	"github.com/google/uuid"
)

type userService struct {
	userRepository ports.UserRepository
}

func NewUserService(repo ports.UserRepository) *userService {
	return &userService{userRepository: repo}
}
func (this *userService) Create(user domain.User) (domain.User, error) {
	// TODO: Check for existing user in database before trying to insert
	user.Id = uuid.New().String()
	user, err := this.userRepository.Insert(user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil

}
func (this *userService) GetAll() ([]domain.User, error) {
	allUsers, err := this.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return allUsers, nil
}
