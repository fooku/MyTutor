package service

import (
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
)

func AddImg(img *models.Img) error {

	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("img")

	err := c.Insert(&img)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "imsert img ไม่ได้"}
	}
	return nil
}

func ListImg() (*[]models.Img, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var img []models.Img

	err := s.DB(models.Database).C("img").Find(nil).Sort("-timestamp").All(&img)
	if err != nil {
		return &img, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return &img, nil
}
