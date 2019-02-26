package api

import (
	"fmt"
	"net/http"

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
