package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

type Date struct {
	Time int `json:"time"`
}

func BuyCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only"}
	}

	d := new(Date)
	err := c.Bind(d)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "time"}
	}
	// fmt.Println("id >", claims["id"])
	idUser := c.FormValue("iduser")
	idCourse := c.FormValue("idcourse")
	fmt.Println("id u>", idUser)
	fmt.Println("id c>", idCourse)
	fmt.Println("date >", d.Time)

	err = service.BuyCourse(idUser, idCourse, d.Time)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
