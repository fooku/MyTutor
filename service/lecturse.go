package service

import (
	"fmt"
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddLectures(lecturse *models.Lectures, id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	idLec := bson.NewObjectId()
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("lectures")

	lecturse.ID = idLec

	var si models.SectionInsert
	err := s.DB(models.Database).C("section").Find(bson.M{"_id": objectID}).One(&si)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	err = c.Insert(&lecturse)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	si.Lectures = append(si.Lectures, idLec)
	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"lectures": si.Lectures}}
	err = s.DB(models.Database).C("section").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func GetLecturesOne(id string) (*models.Lectures, error) {
	s := models.MongoSession.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	Obid := bson.ObjectIdHex(id)
	var lec models.Lectures

	err := s.DB(models.Database).C("lectures").Find(bson.M{"_id": Obid}).One(&lec)
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาlectures one ไม่เจอ"}
	}

	return &lec, nil
}

func UpdateLectures(id string, lectures *models.Lectures) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	fmt.Println(lectures)
	change := bson.M{"$set": bson.M{"name": lectures.Name, "time": lectures.Time, "link": lectures.Link}}
	err := s.DB(models.Database).C("lectures").Update(colQuerier, change)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update lecturesไม่ได้"}
	}

	return nil
}

func DeleteLectures(idlec, idsec string) error {
	fmt.Println("id", idlec)
	if !bson.IsObjectIdHex(idlec) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectIDlec := bson.ObjectIdHex(idlec)
	s := models.MongoSession.Copy()
	defer s.Close()
	err := s.DB(models.Database).C("lectures").RemoveId(objectIDlec)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}
	objectIDsec := bson.ObjectIdHex(idsec)

	var sectionIn models.SectionInsert
	err = s.DB(models.Database).C("section").Find(bson.M{"_id": objectIDsec}).One(&sectionIn)

	for i, v := range sectionIn.Lectures {
		if v == objectIDlec {
			sectionIn.Lectures = append(sectionIn.Lectures[:i], sectionIn.Lectures[i+1:]...)
			fmt.Println("i", i)
		}
	}
	colQuerier := bson.M{"_id": objectIDsec}
	change := bson.M{"$set": bson.M{"lectures": sectionIn.Lectures}}
	err = s.DB(models.Database).C("section").Update(colQuerier, change)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update lecturesไม่ได้"}
	}

	return nil
}
