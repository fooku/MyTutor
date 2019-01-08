package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddPromotion(pr *models.Promotion) error {

	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("promotion")

	pr.Timestamp = time.Now()
	err := c.Insert(&pr)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func ListPromotion() (*[]models.Promotion, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var pr []models.Promotion

	err := s.DB(models.Database).C("promotion").Find(nil).Sort("-timestamp").All(&pr)
	if err != nil {
		return &pr, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return &pr, nil
}

func UpdatePromotion(id string, pr *models.Promotion) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"title": pr.Title, "detail": pr.Detail, "thumbnail": pr.Thumbnail}}
	err := s.DB(models.Database).C("promotion").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func DeletePromotion(id string) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()
	err := s.DB(models.Database).C("promotion").RemoveId(objectID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}
	return nil
}

func GetPromotionOne(id string) (models.Promotion, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	var results models.Promotion
	err := s.DB(models.Database).C("promotion").Find(bson.M{"_id": objectID}).One(&results)

	return results, err
}
