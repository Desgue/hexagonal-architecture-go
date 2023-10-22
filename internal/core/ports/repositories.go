package ports

import "github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"

type UserRepository interface {
	Insert(user domain.User) (domain.User, error)
	FindAll() ([]domain.User, error)
	FindById(id string) (domain.User, error)
}
