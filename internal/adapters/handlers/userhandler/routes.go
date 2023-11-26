package userhandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, handler *UserHttpHandler) {
	r.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})
	v1 := r.Group("/users")
	v1.GET("", handler.GetUsers)
	v1.GET("/id/:id", handler.GetUserById)
	v1.POST("", handler.SaveUser)

}
