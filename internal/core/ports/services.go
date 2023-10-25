package ports

import "github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"

type UserService interface {
	Create(user domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetById(id string) (domain.User, error)
}
