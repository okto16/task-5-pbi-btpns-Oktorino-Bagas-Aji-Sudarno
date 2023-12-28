package router

import (
	"golang-api/controllers"
	"golang-api/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users/register", controllers.RegisterUser)
	r.POST("/users/login", controllers.LoginUser)
	r.GET("/users", controllers.GetUser)

	// Menggunakan middleware
	authorized := r.Group("/users")
	authorized.Use(middlewares.AuthMiddleware())
	{
		authorized.PUT("/:userId", controllers.UpdateUser)
		authorized.DELETE("/:userId", controllers.DeleteUser)
	}
	// Menggunakan middleware
	auth := r.Group("/photos")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/", controllers.CreatePhoto)
		auth.GET("/", controllers.GetPhotos)
		auth.PUT("/:photoId", controllers.UpdatePhoto)
		auth.DELETE("/:photoId", controllers.DeletePhoto)
	}

	return r
}
