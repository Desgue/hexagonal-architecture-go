package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/service/userservice"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/repositories/usersrepo"
	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestGetUsers(t *testing.T) {
	users := map[string]domain.User{
		"1": {Id: "1", Name: "Tester1"},
		"2": {Id: "2", Name: "Tester2"},
	}
	list := []domain.User{{Id: "1", Name: "Tester1"}, {Id: "2", Name: "Tester2"}}
	want, _ := json.Marshal(list)

	r := gin.Default()
	repo := usersrepo.NewFakeRepository(users)
	service := userservice.NewUserService(repo)
	handler := NewUserHttpHandler(service)
	UsersRoute(r, handler)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), string(want))

}

func TestGetUserById(t *testing.T) {
	userMap := map[string]domain.User{"1": {Id: "1", Name: "Tester1"}}
	want, _ := json.Marshal(userMap["1"])

	repo := usersrepo.NewFakeRepository(userMap)
	service := userservice.NewUserService(repo)
	handler := NewUserHttpHandler(service)

	r := gin.Default()
	UsersRoute(r, handler)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/id/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)
	assert.Equal(t, w.Body.String(), string(want))

}
