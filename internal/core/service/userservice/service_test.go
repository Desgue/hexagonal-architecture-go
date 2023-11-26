package userservice

import (
	"reflect"
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/adapters/repositories/usersrepo"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
)

func TestCreate(t *testing.T) {
	user := domain.User{
		Name: "Tester",
	}
	repo := usersrepo.NewFakeRepository(map[string]domain.User{})
	service := NewUserService(repo)

	got, err := service.Create(user)
	if err != nil {
		t.Error(err)
	}

	if reflect.DeepEqual(got.Id, "") {
		t.Error("Error as Id was not successfully created")
	}
	if !reflect.DeepEqual(got.Name, user.Name) {
		t.Errorf("Got %s want %s", got.Name, user.Name)
	}
}

func TestGetAll(t *testing.T) {
	users := []domain.User{
		{Id: "1", Name: "Tester1"},
		{Id: "2", Name: "Tester2"},
	}
	repo := usersrepo.NewFakeRepository(map[string]domain.User{
		"1": {Id: "1", Name: "Tester1"},
		"2": {Id: "2", Name: "Tester2"},
	})

	service := NewUserService(repo)
	got, err := service.GetAll()
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(got, users) {
		t.Errorf("Got %q want %q", got, users)
	}

}
