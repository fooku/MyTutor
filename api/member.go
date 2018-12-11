package api

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
)

// ListMember > ขอรายชื่อสมาชิกทั้งหมด
func ListMember(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if !claims["admin"].(bool) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	u, err := models.GetMember()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
