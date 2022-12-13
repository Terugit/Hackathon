package main

import (
	"db/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	Router()
}

func Router() {
	g := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins: []string{
			"*",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
		},
	}
	g.Use(cors.New(corsConfig))

	//g.POST("/signup", controller.Signup)
	//g.POST("/login", controller.Login)

	api := g.Group("/api")

	api.Use(cors.New(corsConfig))
	//api.Use(middleware.AuthMiddleware())

	api.GET("/member", controller.FetchAllUsers)
	api.POST("invite", controller.CreateUser)
	api.GET("/thanks", controller.FetchAllThank)
	api.GET("/thanks/sent", controller.FetchAllThankSent)
	api.GET("/thanks/received", controller.FetchAllThankReceived)
	//api.GET("/thanks/report", controller.FetchThankReport)
	api.POST("/thanks", controller.CreateThank)
	api.PUT("/thanks", controller.EditThank)
	api.DELETE("/thanks", controller.DeleteThank)

	api.GET("/me", controller.FetchUserInfo)
	//api.PUT("/user", controller.ChangeUserAttributes)
	api.DELETE("/user", controller.DeleteUser)

	if err := g.Run(":8080"); err != nil {
		panic(err)
	}
}
