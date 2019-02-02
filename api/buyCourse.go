package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

func BuyCourse(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only"}
	}

	// fmt.Println("id >", claims["id"])
	idUser := c.FormValue("iduser")
	idCourse := c.FormValue("idcourse")
	fmt.Println("id u>", idUser)
	fmt.Println("id c>", idCourse)

	err := service.BuyCourse(idUser, idCourse)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
