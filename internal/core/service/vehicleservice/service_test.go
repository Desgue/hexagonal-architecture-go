package vehicleservice

import (
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/pkg/apperr"
	"github.com/magiconair/properties/assert"
)

func TestCreateVehicle(t *testing.T) {
	t.Run("Assert if vehicle was created with UUID and inserted in db", func(t *testing.T) {
		repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{}}
		svc := VehicleService(repo)

		newV := domain.Vehicle{Price: 10000}
		got, err := svc.Create(newV)

		if got.ID == "" {
			t.Error("Error creating UUID")
		}
		assert.Equal(t, got.Price, newV.Price)
		assert.Equal(t, err, nil)
	})
	t.Run("Assert vehicle is not inserted if already present in db", func(t *testing.T) {
		repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
			"1": {ID: "1", Price: 1000},
		}}
		svc := VehicleService(repo)

		newV := domain.Vehicle{ID: "1", Price: 1000}
		got, err := svc.Create(newV)

		assert.Equal(t, got, domain.Vehicle{})
		assert.Equal(t, err, apperr.EntryExist)

	})
}

func TestGetAllVehicles(t *testing.T) {
	repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
		"1": {ID: "1", Price: 1000},
		"2": {ID: "2", Price: 2000},
	}}
	svc := VehicleService(repo)

	got, err := svc.GetAll()

	assert.Equal(t, err, nil)
	assert.Equal(t, got, []domain.Vehicle{
		{ID: "1", Price: 1000},
		{ID: "2", Price: 2000},
	})
}

func TestGetVehicleByID(t *testing.T) {
	t.Run("Assert vehicle is present and found in database", func(t *testing.T) {
		repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
			"1": {ID: "1", Price: 1000},
			"2": {ID: "2", Price: 2000},
		}}
		svc := VehicleService(repo)

		got, err := svc.GetById("1")

		assert.Equal(t, err, nil)
		assert.Equal(t, got, domain.Vehicle{ID: "1", Price: 1000})

	})
	t.Run("Assert case vehicle is not present in DB and error is throw", func(t *testing.T) {
		repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
			"1": {ID: "1", Price: 1000},
			"2": {ID: "2", Price: 2000},
		}}
		svc := VehicleService(repo)

		got, err := svc.GetById("3")

		assert.Equal(t, err, apperr.NotFound)
		assert.Equal(t, got, domain.Vehicle{})
	})
}

func TestUpdateVehicle(t *testing.T) {
	t.Run("Assert case vehicle is present and updated in db", func(t *testing.T) {
		repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
			"1": {ID: "1", Price: 1000},
			"2": {ID: "2", Price: 2000},
		}}
		svc := VehicleService(repo)
		want := domain.Vehicle{ID: "1", Price: 5000}
		got, err := svc.Update(want)

		assert.Equal(t, err, nil)
		assert.Equal(t, got, want)

	})
	t.Run("Assert case vehicle is not present in db", func(t *testing.T) {
		repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
			"1": {ID: "1", Price: 1000},
			"2": {ID: "2", Price: 2000},
		}}
		svc := VehicleService(repo)
		want := domain.Vehicle{ID: "3", Price: 3000}
		got, err := svc.Update(want)

		assert.Equal(t, err, apperr.NotFound)
		assert.Equal(t, got, domain.Vehicle{})
	})
}

func TestDeleteVehicle(t *testing.T) {
	repo := fakeRepository{MemoryStorage: map[string]domain.Vehicle{
		"1": {ID: "1", Price: 1000},
		"2": {ID: "2", Price: 2000},
	}}
	svc := VehicleService(repo)
	err := svc.Delete("1")

	assert.Equal(t, err, nil)

}

// Fake Repository Struct and Methods //

type fakeRepository struct {
	MemoryStorage map[string]domain.Vehicle
}

func (fr fakeRepository) Insert(v domain.Vehicle) (domain.Vehicle, error) {
	fr.MemoryStorage[v.ID] = v
	return fr.MemoryStorage[v.ID], nil
}
func (fr fakeRepository) FindAll() ([]domain.Vehicle, error) {
	var vList []domain.Vehicle
	for _, value := range fr.MemoryStorage {
		vList = append(vList, value)
	}
	return vList, nil
}
func (fr fakeRepository) FindById(id string) (domain.Vehicle, error) {
	for _, value := range fr.MemoryStorage {
		if value.ID == id {
			return value, nil
		}
	}
	return domain.Vehicle{}, apperr.NotFound
}
func (fr fakeRepository) Update(v domain.Vehicle) (domain.Vehicle, error) {
	fr.MemoryStorage[v.ID] = v
	return fr.MemoryStorage[v.ID], nil
}
func (fr fakeRepository) Delete(id string) error {
	delete(fr.MemoryStorage, id)
	if _, ok := fr.MemoryStorage[id]; ok {
		return apperr.DeleteErr
	}
	return nil
}
