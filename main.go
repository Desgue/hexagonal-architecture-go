package main

import (
	"github.com/Desgue/hexagonal-architecture-go-example/internal/adapters/handlers/userhandler"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/adapters/repositories/usersrepo"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/service/userservice"
	"github.com/gin-gonic/gin"
)

func main() {
	users := map[string]domain.User{
		"1": {Id: "1", Name: "Tester1"},
		"2": {Id: "2", Name: "Tester2"},
	}

	r := gin.Default()
	repo := usersrepo.NewFakeRepository(users)
	service := userservice.NewUserService(repo)
	httpHandler := userhandler.NewUserHttpHandler(service)
	userhandler.RegisterRoutes(r, httpHandler)

	r.Run()
}
