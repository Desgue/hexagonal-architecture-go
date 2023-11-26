package ports

import "github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"

type UserService interface {
	Create(user domain.User) (domain.User, error)
	GetAll() ([]domain.User, error)
	GetById(id string) (domain.User, error)
}

type VehicleService interface {
	Create(domain.Vehicle) (domain.Vehicle, error)
	GetAll() ([]domain.Vehicle, error)
	GetById(id string) (domain.Vehicle, error)
	Update(domain.Vehicle) (domain.Vehicle, error)
	Delete(id string) error
}
