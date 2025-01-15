package main

import (
	"log"
	"reflect"
	"strings"

	"myproject/database"
	"myproject/routes"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Initialize the database
	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	r := gin.Default()

	// Register custom validation (optional)
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	routes.SetupRoutes(r)
	r.Run("localhost:8080")
}
