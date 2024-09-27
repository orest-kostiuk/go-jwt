package main

import (
	"github.com/gin-gonic/gin"
	"github.com/orest-kostiuk/go-jwt/controllers"
	"github.com/orest-kostiuk/go-jwt/initializers"
	"github.com/orest-kostiuk/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	err := r.Run()
	if err != nil {
		return
	} // listen and serve on 0.0.0.0:8080
}
