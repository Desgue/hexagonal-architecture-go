package userhandler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/service/userservice"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/repositories/usersrepo"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

var (
	data = map[string]domain.User{
		"1": {Id: "1", Name: "Tester1"},
		"2": {Id: "2", Name: "Tester2"},
	}
	repo    = usersrepo.NewFakeRepository(data)
	service = userservice.NewUserService(repo)
	handler = NewUserHttpHandler(service)
	r       *gin.Engine
)

func TestMain(m *testing.M) {
	r = gin.Default()
	RegisterRoutes(r, handler)

	exitVal := m.Run()

	os.Exit(exitVal)

}

func TestGetUsers(t *testing.T) {

	list := []domain.User{data["1"], data["2"]}
	want, _ := json.Marshal(list)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), string(want))

}

func TestGetUserById(t *testing.T) {
	want, _ := json.Marshal(data["1"])

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/id/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), string(want))

}

func TestSaveUser(t *testing.T) {
	want, _ := json.Marshal(domain.User{Id: "3", Name: "Tester3"})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(want))
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), string(want))

}
