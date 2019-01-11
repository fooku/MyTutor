package service

import (
	"fmt"
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddSection(section *models.SectionInsert, id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	idSec := bson.NewObjectId()
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("section")

	section.ID = idSec

	var ci models.CourseInsert
	err := s.DB(models.Database).C("course").Find(bson.M{"_id": objectID}).One(&ci)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	err = c.Insert(&section)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	ci.Section = append(ci.Section, idSec)
	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"section": ci.Section}}
	err = s.DB(models.Database).C("course").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
