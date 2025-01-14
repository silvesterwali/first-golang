package routes

import (
	"myproject/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api:=r.Group("/api/v1")
	{
		users:=api.Group("/users")
		{
			users.POST("/",controllers.NewUserHandler().CreateUser)
			users.GET("/",controllers.NewUserHandler().GetUsers)
			users.GET("/:id",controllers.NewUserHandler().GetUser)
			users.PUT("/:id",controllers.NewUserHandler().UpdateUser)
			users.DELETE("/:id",controllers.NewUserHandler().DeleteUser)
		}

		albums:=api.Group("/albums")
		{
			albums.GET("/",controllers.CreateAlbum)
			albums.GET("/:id",controllers.GetAlbum)
			albums.POST("/",controllers.CreateAlbum)
			albums.PUT("/:id",controllers.UpdateAlbum)
			albums.DELETE("/:id",controllers.DeleteAlbum)
		}

	}
}