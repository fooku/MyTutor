package main

import (
	"log"
	"os"

	"github.com/fooku/LearnOnline_Api/api"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	mongoURL = "mongodb://banana:banana1234@ds123834.mlab.com:23834/video_online"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	err := models.Init(mongoURL)
	if err != nil {
		log.Fatalf("can not init model; %v", err)
	}

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Login route
	e.POST("/login", api.Login)
	e.POST("/register", api.Register)
	// Unauthenticated route
	e.GET("/", api.Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", api.Restricted)
	r.GET("/list", api.ListMember)

	e.Logger.Fatal(e.Start(":" + port))
}
