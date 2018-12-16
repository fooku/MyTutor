package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fooku/LearnOnline_Api/api"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	mongoURL = "mongodb://test:test1234@ds135704.mlab.com:35704/gutututor"
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
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			fmt.Println("Request >>>", c.Request())
		},
	}))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	// Login route
	e.POST("/login", api.Login)
	e.POST("/register", api.Register)
	// Unauthenticated route
	e.GET("/", api.Accessible)

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	// r.GET("", api.Restricted)
	r.GET("/auth", api.Restricted)
	r.GET("/member", api.ListMember)
	r.PUT("/member", api.UpdateMember)
	r.DELETE("/member", api.DeleteMember)

	e.Logger.Fatal(e.Start(":" + port))
}
