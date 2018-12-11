package api

import (
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
)

// ListMember > ขอรายชื่อสมาชิกทั้งหมด
func ListMember(c echo.Context) error {
	u, err := models.GetMember()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}
