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
			// fmt.Println("Request >>>", c.Request())
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
	e.GET("/homecontent", api.ListHomeContent)
	e.GET("/news", api.ListNews)
	e.GET("/promotion", api.ListPromotion)
	e.GET("/course", api.ListCourse)
	e.POST("/lectures", api.AddLectures)
	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	// r.GET("", api.Restricted)
	r.Static("/video", "video")
	r.GET("/auth", api.Restricted)
	r.GET("/member", api.ListMember)
	r.PUT("/member/usertype", api.UpdateMemberUsertype)
	r.PUT("/member", api.UpdateMember)
	r.DELETE("/member", api.DeleteMember)

	r.GET("/homecontentfirstone", api.GetOneFirst)
	r.GET("/homecontentsecondone", api.GetOneSecond)
	r.GET("/homecontentthirdone", api.GetOneThird)

	r.POST("/homecontentfirst", api.AddContenFirst)
	r.POST("/homecontentthird", api.AddContenThird)
	r.PUT("/homecontentfirst", api.UpdateHomeContentFirst)
	r.PUT("/homecontentsecond", api.UpdateHomeContentSecond)
	r.PUT("/homecontentthird", api.UpdateHomeContentThird)
	r.DELETE("/homecontentfirst", api.DeleteHomeContentFirst)
	r.DELETE("/homecontentthird", api.DeleteHomeContentThird)

	r.GET("/newsone", api.GetNews)
	r.POST("/news", api.AddNews)
	r.PUT("/news", api.UpdateNews)
	r.DELETE("/news", api.DeleteNews)

	r.GET("/promotionone", api.GetPromotion)
	r.POST("/promotion", api.AddPromotion)
	r.PUT("/promotion", api.UpdatePromotion)
	r.DELETE("/promotion", api.DeletePromotion)

	r.GET("/courseall", api.ListCourseAll)
	r.GET("/courseone", api.ListCourseOne)
	r.POST("/course", api.AddCourse)
	r.PUT("/course", api.UpdateCourse)
	r.DELETE("/course", api.DeleteCourse)

	r.GET("/sectionone", api.GetSectionOne)
	r.POST("/section", api.AddSection)
	r.PUT("/section", api.UpdateSection)
	r.DELETE("/section", api.DeleteSection)

	r.GET("/lecturesone", api.ListLecturesOne)
	r.POST("/lectures", api.AddLectures)
	r.PUT("/lectures", api.UpdateLectures)
	r.DELETE("/lectures", api.DeleteLectures)

	e.Logger.Fatal(e.Start(":" + port))
}
