package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func BuyCourse(idUser, idCourse string, date int) error {
	if !bson.IsObjectIdHex(idUser) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	if !bson.IsObjectIdHex(idCourse) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}

	objectIDUser := bson.ObjectIdHex(idUser)
	objectIDCourse := bson.ObjectIdHex(idCourse)
	s := models.MongoSession.Copy()
	defer s.Close()

	var user models.User
	colQuerierUser := bson.M{"_id": objectIDUser}
	err := s.DB(models.Database).C("users").Find(bson.M{"_id": objectIDUser}).One(&user)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาUser one ไม่ได้"}
	}
	var cci models.CourseInsert
	colQuerier := bson.M{"_id": objectIDCourse}
	err = s.DB(models.Database).C("course").Find(bson.M{"_id": objectIDCourse}).One(&cci)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาCourse one ไม่ได้"}
	}

	var cu models.ClaimCourse
	cu.ID = bson.NewObjectId()
	cu.IDCourse = objectIDCourse
	cu.IDUser = objectIDUser
	cu.TimeStart = time.Now()
	cu.TimeEnd = time.Now().AddDate(0, 0, date)
	err = s.DB(models.Database).C("claimcourse").Insert(&cu)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Insert claimcourse ไม่ได้"}
	}

	cci.ClaimUser = append(cci.ClaimUser, cu.ID)
	change := bson.M{"$set": bson.M{"claimuser": cci.ClaimUser}}
	err = s.DB(models.Database).C("course").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	user.MyCourse = append(user.MyCourse, cu.ID)
	changeUser := bson.M{"$set": bson.M{"mycourse": user.MyCourse}}
	err = s.DB(models.Database).C("users").Update(colQuerierUser, changeUser)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update mycourse in course ไม่ได้"}
	}

	return nil
}
