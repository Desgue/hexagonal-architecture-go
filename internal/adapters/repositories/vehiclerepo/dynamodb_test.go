package vehiclerepo

import (
	"context"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
)

var (
	testRepo *dynamoRepository
)

func TestMain(m *testing.M) {
	log.Println("Setup before all tests")

	ctx := context.Background()
	localStackCont, err := prepareContainer(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	testRepo = NewDynamoRepository(localStackCont.URI, "vehicles")

	exitVal := m.Run()

	log.Println("After all tests")

	if err = localStackCont.Terminate(ctx); err != nil {
		log.Fatalln(err)
	}
	os.Exit(exitVal)
}

func TestInsert(t *testing.T) {
	log.Println("TestInsert")
	if err := createTable(testRepo); err != nil {
		log.Fatalln(err)
	}
	defer deleteTable(testRepo)
	newVehicle := domain.Vehicle{
		ID:    "1",
		Price: 100,
	}
	gotVehicle, err := testRepo.Insert(newVehicle)
	if err != nil {
		t.Errorf("Got error: %s when inserting new user", err)
	}

	if !reflect.DeepEqual(gotVehicle.ID, "1") {
		t.Errorf("Got %v want %v", gotVehicle.ID, "1")
	}
	if !reflect.DeepEqual(gotVehicle.Price, 100) {
		t.Errorf("Got %v want %v", gotVehicle.Price, 100)
	}

}

func TestFindAll(t *testing.T) {
	log.Println("Test FindAll")
	if err := createTable(testRepo); err != nil {
		log.Fatalln(err)
	}
	defer deleteTable(testRepo)
	want, err := populateDb(*testRepo)
	if err != nil {
		t.Errorf("Got error populating database: %s", err)
	}

	got, err := testRepo.FindAll()
	if err != nil {
		t.Errorf("Error fetching results from db %s", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, want %v", got, want)
	}
}
func TestFindById(t *testing.T) {
	log.Println("Test FindById")
	if err := createTable(testRepo); err != nil {
		log.Fatalln(err)
	}
	defer deleteTable(testRepo)
	_, err := populateDb(*testRepo)
	if err != nil {
		t.Errorf("Got error populating database: %s", err)
	}
	gotVehicle, err := testRepo.FindById("1")
	if err != nil {
		t.Errorf("Got error searching user: %s", err)
	}
	if !reflect.DeepEqual(gotVehicle.ID, "1") {
		t.Errorf("Got %v want %v", gotVehicle.ID, "1")
	}
	if !reflect.DeepEqual(gotVehicle.Price, 100) {
		t.Errorf("Got %v want %v", gotVehicle.Price, 100)
	}

}
