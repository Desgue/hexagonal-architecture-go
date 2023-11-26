package vehicleservice

import (
	"fmt"
	"log"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/ports"
	"github.com/Desgue/hexagonal-architecture-go-example/pkg/apperr"
	"github.com/google/uuid"
)

type vehicleService struct {
	repo ports.VehicleRepository
}

func VehicleService(repo ports.VehicleRepository) vehicleService {
	return vehicleService{
		repo: repo,
	}
}

func (vs vehicleService) Create(v domain.Vehicle) (domain.Vehicle, error) {
	// check if vehicle is not present in database
	_, err := vs.repo.FindById(v.ID)
	switch err {
	case apperr.NotFound:
		// if not present create UUID and insert in database
		v.ID = uuid.New().String()
		insertedV, err := vs.repo.Insert(v)
		if err != nil {
			return domain.Vehicle{}, fmt.Errorf("Error creating vehicle in database from service: %s", err)
		}
		return insertedV, nil
	case nil:
		// if present return error indicating that
		return domain.Vehicle{}, apperr.EntryExist
	default:
		// if any other error  log and return err
		log.Printf("Error finding vehicle by ID in database from service: %s", err)
		return domain.Vehicle{}, err

	}

}

func (vs vehicleService) GetAll() ([]domain.Vehicle, error) {
	var allV []domain.Vehicle
	allV, err := vs.repo.FindAll()
	if err != nil {
		return []domain.Vehicle{}, fmt.Errorf("Error getting all vehicles from service: %s", err)
	}
	return allV, nil

}

func (vs vehicleService) GetById(id string) (domain.Vehicle, error) {
	v, err := vs.repo.FindById(id)
	if err != nil {
		return domain.Vehicle{}, err
	}
	return v, nil

}

func (vs vehicleService) Update(v domain.Vehicle) (domain.Vehicle, error) {
	// check if vehicle is present in db
	_, err := vs.repo.FindById(v.ID)
	if err != nil {
		return domain.Vehicle{}, err
	}
	// if present update
	updatedV, err := vs.repo.Insert(v)
	if err != nil {
		return domain.Vehicle{}, err
	}
	return updatedV, nil

}

func (vs vehicleService) Delete(id string) error {
	err := vs.repo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
