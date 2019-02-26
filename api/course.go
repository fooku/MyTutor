package api

import (
	"fmt"
	"net/http"
	"strconv"

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

func UpdateCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	course := new(models.CourseInsert)
	err := c.Bind(course)
	fmt.Println(id)
	fmt.Println("c >>", course.Type)
	fmt.Println()

	err = service.UpdateCourse(id, course)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func DeleteCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")

	err := service.DeleteCourse(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func UpdatePublishCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	p, err := strconv.ParseBool(c.FormValue("p"))
	fmt.Println(id, p)
	if err != nil {
		return err
	}
	err = service.UpdatePublishCourse(id, p)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func GetCourseOnePublish(c echo.Context) error {
	id := c.FormValue("id")

	course, err := service.GetCourseOnePublish(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, course)
}

func GetMyCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	id := c.FormValue("id")

	idc := claims["id"].(string)
	course, err := service.GetMyCourse(id, idc)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, course)
}
