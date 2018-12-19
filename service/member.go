package service

import (
	"fmt"
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func GetMember() ([]models.User, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var results []models.User
	err := s.DB(models.Database).C("users").Find(nil).Sort("-timestamp").All(&results)

	fmt.Println("Results All: ", results)
	return results, err
}

func GetMemberOne(id string) (models.User, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	var results models.User
	err := s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&results)

	return results, err
}

func DeleteMember(id string) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()
	err := s.DB(models.Database).C("users").RemoveId(objectID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}
	return nil
}

func UpdateUserType(id string, usertype string) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"usertype": usertype}}
	err := s.DB(models.Database).C("users").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
