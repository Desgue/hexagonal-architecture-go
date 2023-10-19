package usersrepo

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/localstack"
)

func TestInsert(t *testing.T) {
	ctx := context.Background()
	localStackCont, err := prepareContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := localStackCont.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	repo := NewDynamoRepository(localStackCont.URI)
	if err := repo.createTable(); err != nil {
		t.Fatal(err)
	}
	newUser := domain.User{
		Id:   "1",
		Name: "Tester",
	}
	gotUser, err := repo.Insert(newUser)
	if err != nil {
		t.Errorf("Got error: %s when inserting new user", err)
	}

	if !reflect.DeepEqual(gotUser.Id, "1") {
		t.Errorf("Got %v want %v", gotUser.Id, "1")
	}
	if !reflect.DeepEqual(gotUser.Name, "Tester") {
		t.Errorf("Got %v want %v", gotUser.Name, "Tester")
	}

}
func TestFindById(t *testing.T) {
	ctx := context.Background()
	localStackCont, err := prepareContainer(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := localStackCont.Terminate(ctx); err != nil {
			panic(err)
		}
	}()
	repo := NewDynamoRepository(localStackCont.URI)
	if err := repo.createTable(); err != nil {
		t.Fatal(err)
	}
	newUser := domain.User{
		Id:   "1",
		Name: "Tester",
	}
	_, err = repo.Insert(newUser)
	if err != nil {
		t.Errorf("Got error inserting user: %s", err)
	}
	gotUser, err := repo.FindById(newUser.Id)
	if err != nil {
		t.Errorf("Got error searching user: %s", err)
	}
	if !reflect.DeepEqual(gotUser.Id, "1") {
		t.Errorf("Got %v want %v", gotUser.Id, "1")
	}
	if !reflect.DeepEqual(gotUser.Name, "Tester") {
		t.Errorf("Got %v want %v", gotUser.Name, "Tester")
	}

}

type container struct {
	*localstack.LocalStackContainer
	URI string
}

func prepareContainer(ctx context.Context) (*container, error) {
	cont, err := localstack.RunContainer(ctx, testcontainers.WithImage("localstack/localstack:latest"))
	if err != nil {
		return nil, err
	}
	provider, err := testcontainers.NewDockerProvider()
	if err != nil {
		return nil, err
	}
	host, err := provider.DaemonHost(ctx)
	if err != nil {
		return nil, err
	}
	mappedPort, err := cont.MappedPort(ctx, nat.Port("4566/tcp"))
	if err != nil {
		return nil, err
	}
	URI := fmt.Sprintf("http://%s:%d", host, mappedPort.Int())
	return &container{LocalStackContainer: cont, URI: URI}, nil
}
