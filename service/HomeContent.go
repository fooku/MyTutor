package service

import (
	"fmt"
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddHomeContent(ahc models.AddHomeContent) error {
	err := ahc.AddContent()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func UpdateHomeContent(ahc models.UpdateHomeContent, id string) error {
	err := ahc.UpdateContent(id)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func DeleteHomeContent(ahc models.AddHomeContent, id string) error {
	err := ahc.DeleteContent(id)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func GetHomeContent() (*models.HomeContent, error) {
	content := new(models.HomeContent)
	s := models.MongoSession.Copy()
	defer s.Close()

	var first []models.HomeContentFirst
	var second []models.HomeContentSecond
	var third []models.HomeContentThird
	err := s.DB(models.Database).C("homecontentfirst").Find(nil).All(&first)
	if err != nil {
		return content, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	err = s.DB(models.Database).C("homecontentsecond").Find(nil).All(&second)
	if err != nil {
		return content, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	err = s.DB(models.Database).C("homecontentthird").Find(nil).All(&third)
	if err != nil {
		return content, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	content.HomeContentFirst = first
	content.HomeContentSecond = second
	content.HomeContentThird = third

	fmt.Println("Results HomeContent: ", content)
	return content, err
}

func GetOneHomeContentF(id string) (models.HomeContentFirst, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	fmt.Println("objectID", objectID)
	var results models.HomeContentFirst
	err := s.DB(models.Database).C("homecontentfirst").Find(bson.M{"_id": objectID}).One(&results)

	return results, err
}

func GetOneHomeContentS(id string) (models.HomeContentSecond, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	var results models.HomeContentSecond
	err := s.DB(models.Database).C("homecontentsecond").Find(bson.M{"_id": objectID}).One(&results)

	return results, err
}

func GetOneHomeContentT(id string) (models.HomeContentThird, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	var results models.HomeContentThird
	err := s.DB(models.Database).C("homecontentthird").Find(bson.M{"_id": objectID}).One(&results)

	return results, err
}
