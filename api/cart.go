package api

import (
	"fmt"
	"io"
	"net/http"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

type indexRequest struct {
	Index int `json:"index" `
}

func AddItem(c echo.Context) error {

	iduser := c.FormValue("iduser")
	item := new(models.Item)
	err := c.Bind(item)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "item error"}
	}
	fmt.Println("id user>", iduser)
	fmt.Println("id error>", item)

	err = service.AddCart(iduser, item)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func DeleteItem(c echo.Context) error {
	iduser := c.FormValue("iduser")
	index := new(indexRequest)
	err := c.Bind(index)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "item error"}
	}
	fmt.Println("id user>", iduser)

	err = service.DeleteItem(iduser, index.Index)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func GenOrder(c echo.Context) error {
	iduser := c.FormValue("iduser")

	fmt.Println("id user>", iduser)

	err := service.GenOrder(iduser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
func GetOrderOne(c echo.Context) error {
	idorder := c.FormValue("idorder")

	fmt.Println("id order>", idorder)

	order, err := service.GetOrderOne(idorder)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, order)
}

func ListOrder(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}

	order, err := service.ListOrder()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, order)
}

func AddPayment(c echo.Context) error {
	idorder := c.FormValue("idorder")

	fmt.Println("id order>", idorder)

	payment := new(models.Payment)
	err := c.Bind(payment)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "item error"}
	}
	fmt.Println("id order>", idorder)
	fmt.Println("id error>", payment)

	err = service.AddPayment(idorder, payment)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func AddPayment2(c echo.Context) error {
	idorder := c.FormValue("idorder")
	fmt.Println("id order>", idorder)

	file, err := c.FormFile("file")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "uplode error"}
	}

	img := "/imagepayment/" + file.Filename

	src, err := file.Open()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "uplode error"}
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./imagepayment/" + file.Filename)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "uplode error"}
	}
	defer dst.Close()

	// Copy
	io.Copy(dst, src)

	err = service.AddPayment2(idorder, img)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func BuyMany(c echo.Context) error {
	idorder := c.FormValue("idorder")

	fmt.Println("id order>", idorder)

	fmt.Println("id order>", idorder)

	err := service.BuyMany(idorder, 30)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
