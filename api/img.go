package api

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddImg(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}
	img := new(models.Img)

	img.ID = bson.NewObjectId()

	file, err := c.FormFile("file")
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "uplode error"}
	}

	name := strings.Split(file.Filename, ".")
	img.Name = name[0]
	img.Img = "/img/" + img.ID.Hex() + "." + name[1]
	img.Timestamp = time.Now()

	file.Filename = img.ID.Hex() + "." + name[1]

	src, err := file.Open()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "uplode error"}
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("./img/" + file.Filename)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "uplode error"}
	}
	defer dst.Close()

	// Copy
	io.Copy(dst, src)

	err = service.AddImg(img)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func ListImg(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}

	news, err := service.ListImg()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "err"}
	}
	return c.JSON(http.StatusOK, news)
}
