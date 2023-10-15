package ports

import "github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"

type UserRepository interface {
	Insert() (domain.User, error)
	FindAll() ([]domain.User, error)
}
