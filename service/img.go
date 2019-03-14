package service

import (
	"net/http"
	"os"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
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

func DeleteImg(id string) error {
	s := models.MongoSession.Copy()
	defer s.Close()

	objectID := bson.ObjectIdHex(id)

	var img models.Img
	err := s.DB(models.Database).C("img").Find(bson.M{"_id": objectID}).One(&img)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาlectures one ไม่เจอ"}
	}

	err = s.DB(models.Database).C("img").RemoveId(objectID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}

	err = os.Remove("." + img.Img)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "ลบไม่ได้"}
	}

	return nil
}
