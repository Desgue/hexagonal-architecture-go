package userhandler

import (
	"net/http"

	"github.com/Desgue/hexagonal-architecture-go-example/internal/core/domain"
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
	var user domain.User
	c.BindJSON(&user)
	user, err := this.Service.Create(user)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, user)

}
func (this *UserHttpHandler) GetUsers(c *gin.Context) {
	users, err := this.Service.GetAll()
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, users)

}
func (this *UserHttpHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := this.Service.GetById(id)
	if err != nil {
		c.String(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
