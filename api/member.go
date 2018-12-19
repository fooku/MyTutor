package api

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
)

// ListMember > ขอรายชื่อสมาชิกทั้งหมด
func ListMember(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	u, err := service.GetMember()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

// DeleteMember > ลบสมาชิก
func DeleteMember(c echo.Context) error {
	id := c.FormValue("id")
	fmt.Println(id)
	err := service.DeleteMember(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

//UpdateMember > แก้ไข User
func UpdateMember(c echo.Context) error {
	id := c.FormValue("id")
	u := new(models.UpdateRequest)
	err := c.Bind(u)
	fmt.Println(id)
	err = service.UpdateUserType(id, u.Usertype)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}
