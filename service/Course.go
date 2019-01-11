package service

import (
	"fmt"
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddCourse(course *models.CourseInsert) error {
	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("course")

	err := c.Insert(&course)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	// for _, sec := range course.Section {
	// 	lecc := make([]bson.ObjectId, 0)
	// 	for _, lec := range sec.Lectures {
	// 		id := bson.NewObjectId()
	// 		lec.ID = id
	// 		lecc = append(lecc, id)
	// 		c := s.DB(models.Database).C("lectures")
	// 		err := c.Insert(&lec)
	// 		if err != nil {
	// 			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	// 		}
	// 	}
	// 	c := s.DB(models.Database).C("section")
	// 	secIn := models.SectionInsert{Name: sec.Name, Lectures: lecc}
	// 	err := c.Insert(&secIn)
	// 	if err != nil {
	// 		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	// 	}
	// }

	return nil
}

func ListCourse() (*[]models.Course, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var cci []models.CourseInsert

	err := s.DB(models.Database).C("course").Find(nil).All(&cci)
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	courses := make([]models.Course, len(cci))
	for i, course := range cci {
		var user models.User
		objectID := course.Creator
		err = s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&user)
		if err != nil {
			return &courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
		}

		courses[i].ID = course.ID
		courses[i].Name = course.Name
		courses[i].Hour = course.Hour
		courses[i].Creator = user
		courses[i].Price = course.Price

		section := make([]models.Section, len(course.Section))
		for j, sec := range course.Section {
			si := new(models.SectionInsert)
			err = s.DB(models.Database).C("section").Find(bson.M{"_id": sec}).One(&si)
			if err != nil {
				return &courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
			}
			fmt.Println(sec)
			section[j].ID = si.ID
			section[j].Name = si.Name
			lectures := make([]models.Lectures, len(si.Lectures))
			for index, l := range si.Lectures {
				lec := new(models.Lectures)
				err = s.DB(models.Database).C("lectures").Find(bson.M{"_id": l}).One(&lec)
				if err != nil {
					return &courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
				}
				lectures[index] = *lec
			}
			section[j].Lectures = lectures
		}
		fmt.Println("sec > ", section)
		courses[i].Section = section
		if err != nil {
			return &courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
		}
	}

	return &courses, nil
}
