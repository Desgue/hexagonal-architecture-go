package handler

import (
	"log"
	"net/http"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/ports"
	"github.com/gin-gonic/gin"
)

type UserHttpHandler struct {
	Service ports.UserService
}

func NewUserHttpHandler(service ports.UserService) *UserHttpHandler {
	return &UserHttpHandler{
		Service: service,
	}
}

func (this *UserHttpHandler) SaveUser(c *gin.Context) {

}
func (this *UserHttpHandler) GetUsers(c *gin.Context) {
	users, err := this.Service.GetAll()
	if err != nil {
		c.JSON(404, users)
	}
	if err != nil {
		log.Fatalln(err)
	}
	c.JSON(http.StatusOK, users)

}
func (this *UserHttpHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := this.Service.GetById(id)
	if err != nil {
		c.String(http.StatusNotFound, "User not found")
	}
	c.JSON(http.StatusOK, user)
}
