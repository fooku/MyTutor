package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddNews(news *models.News) error {

	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("news")

	news.Timestamp = time.Now()
	err := c.Insert(&news)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func ListNews() (*[]models.News, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var news []models.News

	err := s.DB(models.Database).C("news").Find(nil).Sort("-timestamp").All(&news)
	if err != nil {
		return &news, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return &news, nil
}

func UpdateNews(id string, news *models.News) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"title": news.Title, "detail": news.Detail, "thumbnail": news.Thumbnail}}
	err := s.DB(models.Database).C("news").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func DeleteNews(id string) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()
	err := s.DB(models.Database).C("news").RemoveId(objectID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}
	return nil
}

func GetNewsOne(id string) (models.News, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	var results models.News
	err := s.DB(models.Database).C("news").Find(bson.M{"_id": objectID}).One(&results)

	return results, err
}
