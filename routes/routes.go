package routes

import (
	"myproject/controllers"
	"myproject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", controllers.NewUserHandler().CreateUser)
			users.GET("/", controllers.NewUserHandler().GetUsers)
			users.GET("/:id", controllers.NewUserHandler().GetUser)
			users.PUT("/:id", controllers.NewUserHandler().UpdateUser)
			users.DELETE("/:id", controllers.NewUserHandler().DeleteUser)
		}

		albums := api.Group("/albums")
		{
			albums.GET("/", controllers.NewAlbumHandler().GetAlbums)
			albums.GET("/:id", controllers.NewAlbumHandler().GetAlbum)
			albums.POST("/", controllers.NewAlbumHandler().CreateAlbum)
			albums.PUT("/:id", controllers.NewAlbumHandler().UpdateAlbum)
			albums.DELETE("/:id", controllers.NewAlbumHandler().DeleteAlbum)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/login", controllers.NewAuthHandler().Login)
			auth.Use(middleware.Auth())
			auth.GET("/profile", controllers.NewAuthHandler().Profile)
		}

	}
}
