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

func GetSectionOne(id string) (*models.Section, error) {
	s := models.MongoSession.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	idSection := bson.ObjectIdHex(id)
	var sectionIn models.SectionInsert

	err := s.DB(models.Database).C("section").Find(bson.M{"_id": idSection}).One(&sectionIn)
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	section := new(models.Section)

	section.ID = sectionIn.ID
	section.Name = sectionIn.Name

	return section, nil
}
func UpdateSection(id string, section *models.SectionInsert) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"name": section.Name}}
	err := s.DB(models.Database).C("section").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func DeleteSection(idsec, idcorse string) error {
	fmt.Println("id", idsec)
	if !bson.IsObjectIdHex(idsec) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	s := models.MongoSession.Copy()
	defer s.Close()

	objectIDsec := bson.ObjectIdHex(idsec)

	var sectionIn models.SectionInsert

	err := s.DB(models.Database).C("section").Find(bson.M{"_id": objectIDsec}).One(&sectionIn)

	for _, lectures := range sectionIn.Lectures {
		err = s.DB(models.Database).C("lectures").RemoveId(lectures)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
		}
	}

	err = s.DB(models.Database).C("section").RemoveId(objectIDsec)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}

	objectIDcourse := bson.ObjectIdHex(idcorse)

	var courseIn models.CourseInsert
	err = s.DB(models.Database).C("course").Find(bson.M{"_id": objectIDcourse}).One(&courseIn)

	fmt.Println("ก่อน : ", courseIn.Section)
	for i, section := range courseIn.Section {
		if section == objectIDsec {
			courseIn.Section = append(courseIn.Section[:i], courseIn.Section[i+1:]...)
			fmt.Println("i", i)
		}
	}
	fmt.Println("หลัง : ", courseIn.Section)
	colQuerier := bson.M{"_id": objectIDcourse}
	change := bson.M{"$set": bson.M{"section": courseIn.Section}}
	err = s.DB(models.Database).C("course").Update(colQuerier, change)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update courseไม่ได้"}
	}

	return nil
}
