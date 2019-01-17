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
		courses[i].Thumbnail = course.Thumbnail
		courses[i].Detail = course.Detail
		courses[i].Type = course.Type

		section := make([]models.Section, len(course.Section))
		fmt.Println(len(course.Section))
		fmt.Println(course.Section)
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

func GetCourse() (*[]models.Course, error) {
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
		courses[i].Thumbnail = course.Thumbnail
		courses[i].Detail = course.Detail
		courses[i].Type = course.Type
	}

	return &courses, nil
}

func GetCourseOne(id string) (*models.Course, error) {
	s := models.MongoSession.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(id) {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	idCci := bson.ObjectIdHex(id)
	var cci models.CourseInsert

	err := s.DB(models.Database).C("course").Find(bson.M{"_id": idCci}).One(&cci)
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาCourse one ไม่ได้"}
	}
	courses := new(models.Course)

	var user models.User
	objectID := cci.Creator
	err = s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&user)
	if err != nil {
		return courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาUser one ไม่ได้"}
	}

	courses.ID = cci.ID
	courses.Name = cci.Name
	courses.Hour = cci.Hour
	courses.Creator = user
	courses.Price = cci.Price
	courses.Thumbnail = cci.Thumbnail
	courses.Detail = cci.Detail
	courses.Type = cci.Type

	section := make([]models.Section, len(cci.Section))
	fmt.Println(len(cci.Section))
	fmt.Println("sec ก่อน", cci.Section)
	for j, sec := range cci.Section {
		si := new(models.SectionInsert)
		err = s.DB(models.Database).C("section").Find(bson.M{"_id": sec}).One(&si)
		if err != nil {
			return courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
		}
		fmt.Println(sec)
		section[j].ID = si.ID
		section[j].Name = si.Name
		lectures := make([]models.Lectures, len(si.Lectures))
		for index, l := range si.Lectures {
			lec := new(models.Lectures)
			err = s.DB(models.Database).C("lectures").Find(bson.M{"_id": l}).One(&lec)
			if err != nil {
				return courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
			}
			lectures[index] = *lec
		}
		section[j].Lectures = lectures
	}
	fmt.Println("sec > ", section)
	courses.Section = section
	if err != nil {
		return courses, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}

	return courses, nil
}

func UpdateCourse(id string, course *models.CourseInsert) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"name": course.Name, "hour": course.Hour,
		"price": course.Price, "type": course.Type, "detail": course.Detail,
		"thumbnail": course.Thumbnail}}
	err := s.DB(models.Database).C("course").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func DeleteCourse(id string) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	s := models.MongoSession.Copy()
	defer s.Close()

	objectIDcourse := bson.ObjectIdHex(id)

	var courseIn models.CourseInsert

	err := s.DB(models.Database).C("course").Find(bson.M{"_id": objectIDcourse}).One(&courseIn)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found course"}
	}

	for _, section := range courseIn.Section {
		var sectionIn models.SectionInsert
		err = s.DB(models.Database).C("section").Find(bson.M{"_id": section}).One(&sectionIn)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found section"}
		}
		fmt.Println("[]lec : ", sectionIn.Lectures)
		for _, lec := range sectionIn.Lectures {
			fmt.Println("lec", lec)
			err = s.DB(models.Database).C("lectures").RemoveId(lec)
			if err != nil {
				return &echo.HTTPError{Code: http.StatusNotFound, Message: "not RemoveId lec"}
			}
		}

		err = s.DB(models.Database).C("section").RemoveId(section)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusNotFound, Message: "not RemoveId section"}
		}
	}

	err = s.DB(models.Database).C("course").RemoveId(objectIDcourse)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not RemoveId course"}
	}

	return nil
}
