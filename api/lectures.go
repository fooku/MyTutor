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

	id := c.FormValue("id")
	lestures := new(models.Lectures)
	err := c.Bind(lestures)
	fmt.Println()
	fmt.Println("input >>>> ", lestures)
	fmt.Println()
	if err != nil {
		return err
	}

	err = service.AddLectures(lestures, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
