package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kweku-xvi/go-users/api/v1/controllers"
	"github.com/kweku-xvi/go-users/api/v1/initializers"
	"github.com/kweku-xvi/go-users/api/v1/middleware"
)

func init() {
	initializers.LoadDotEnv()
	initializers.ConnectToDB()
	initializers.SyncDB()
}

func main() {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	
	r.Run()
}
