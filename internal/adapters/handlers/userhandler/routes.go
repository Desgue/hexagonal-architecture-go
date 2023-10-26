package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *UserHttpHandler) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	v1 := r.Group("/api/v1")
	v1.GET("/users", handler.GetUsers)
	v1.GET("/users/id/:id", handler.GetUserById)
	v1.POST("/users", handler.SaveUser)

}
