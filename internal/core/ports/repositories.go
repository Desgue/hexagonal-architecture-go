package ports

import "github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"

type UserRepository interface {
	Insert(user domain.User) (domain.User, error)
	FindAll() ([]domain.User, error)
	FindById(id string) (domain.User, error)
}

type VehicleRepository interface {
	Insert(v domain.Vehicle) (domain.Vehicle, error)
	FindAll() ([]domain.Vehicle, error)
	FindById(id string) (domain.Vehicle, error)
	Delete(id string) error
}
