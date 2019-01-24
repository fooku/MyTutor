package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

func AddLectures(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}

	idsec := c.FormValue("idsec")
	quality := c.FormValue("quality")

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	fmt.Println()
	fmt.Println("sec >>>> ", idsec)
	fmt.Println()
	if err != nil {
		return err
	}

	err = service.AddLectures(idsec, quality, file)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func ListLecturesOne(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")

	lec, err := service.GetLecturesOne(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, lec)
}

func UpdateLectures(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	lectures := new(models.Lectures)
	err := c.Bind(lectures)
	fmt.Println(id)
	fmt.Println(lectures.Name)
	err = service.UpdateLectures(id, lectures)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func DeleteLectures(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	idsec := c.FormValue("idsec")
	idlec := c.FormValue("idlec")

	err := service.DeleteLectures(idlec, idsec)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
