package api

import (
	"fmt"
	"net/http"

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

func ListHomeContent(c echo.Context) error {
	hc, err := service.GetHomeContent()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, hc)

}
