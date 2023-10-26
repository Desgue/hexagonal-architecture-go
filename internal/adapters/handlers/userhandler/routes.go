package userhandler

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *UserHttpHandler) {
	r.GET("/", handler.Root)
	r.GET("/users", handler.GetUsers)
	r.GET("/users/id/:id", handler.GetUserById)
	r.POST("/users", handler.SaveUser)

}
