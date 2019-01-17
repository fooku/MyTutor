package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}

	course := new(models.CourseInsert)
	err := c.Bind(course)
	fmt.Println()
	fmt.Println("input >>>> ", course)
	fmt.Println()
	if err != nil {
		return err
	}
	course.Creator = bson.ObjectIdHex(claims["id"].(string))

	err = service.AddCourse(course)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func ListCourseAll(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	course, err := service.ListCourse()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, course)
}

func ListCourse(c echo.Context) error {
	course, err := service.GetCourse()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, course)
}

func ListCourseOne(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")

	course, err := service.GetCourseOne(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, course)
}
