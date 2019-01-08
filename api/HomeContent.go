package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

func AddContenFirst(c echo.Context) error {
	hcfr := new(models.HomeContentFirst)
	err := c.Bind(hcfr)
	fmt.Println(hcfr)
	if err != nil {
		return err
	}

	err = service.AddHomeContent(hcfr)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func AddContenThird(c echo.Context) error {
	hcfr := new(models.HomeContentThird)
	err := c.Bind(hcfr)
	fmt.Println(hcfr)
	if err != nil {
		return err
	}

	err = service.AddHomeContent(hcfr)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func UpdateHomeContentFirst(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	hcfr := new(models.HomeContentFirst)
	err := c.Bind(hcfr)
	fmt.Println(hcfr)
	if err != nil {
		return err
	}

	err = service.UpdateHomeContent(hcfr, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
func UpdateHomeContentSecond(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	hs := new(models.HomeContentSecond)
	err := c.Bind(hs)
	if err != nil {
		return err
	}

	err = service.UpdateHomeContent(hs, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func UpdateHomeContentThird(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	h := new(models.HomeContentThird)
	err := c.Bind(h)
	if err != nil {
		return err
	}

	err = service.UpdateHomeContent(h, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func ListHomeContent(c echo.Context) error {
	hc, err := service.GetHomeContent()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, hc)
}

func DeleteHomeContentFirst(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	hcfr := new(models.HomeContentFirst)

	err := service.DeleteHomeContent(hcfr, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func DeleteHomeContentThird(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	hcfr := new(models.HomeContentThird)

	err := service.DeleteHomeContent(hcfr, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func GetOneFirst(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	fmt.Println()
	fmt.Println("id", id)
	fmt.Println()

	n, err := service.GetOneHomeContentF(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, n)
}
func GetOneSecond(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	fmt.Println(id)

	n, err := service.GetOneHomeContentS(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, n)
}
func GetOneThird(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	id := c.FormValue("id")
	fmt.Println(id)

	n, err := service.GetOneHomeContentT(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, n)
}
