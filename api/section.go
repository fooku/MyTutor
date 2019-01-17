package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

func AddSection(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}

	id := c.FormValue("id")
	section := new(models.SectionInsert)
	err := c.Bind(section)
	fmt.Println()
	fmt.Println("input >>>> ", section)
	fmt.Println()
	if err != nil {
		return err
	}

	err = service.AddSection(section, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func GetSectionOne(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")

	section, err := service.GetSectionOne(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, section)
}

func UpdateSection(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	section := new(models.SectionInsert)
	err := c.Bind(section)
	fmt.Println(id)
	err = service.UpdateSection(id, section)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
