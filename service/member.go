package service

import (
	"fmt"
	"net/http"
	"time"

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

func GetMemberOne(id string) (models.UserResponse, error) {
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	var results models.User
	var user models.UserResponse
	err := s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&results)
	if err != nil {
		return user, &echo.HTTPError{Code: http.StatusNotFound, Message: "หา users one ไม่เจอ"}
	}

	user.ID = results.ID
	user.Email = results.Email
	user.FirstName = results.FirstName
	user.LastName = results.LastName
	user.NickName = results.NickName
	user.TelephoneNumber = results.TelephoneNumber
	user.Timestamp = results.Timestamp
	user.UserType = results.UserType
	user.Address = results.Address
	user.Birthday = results.Birthday
	user.Cart = results.Cart
	fmt.Println(">>>>>>>>>>>>", results.Cart)
	var cc []models.ClaimCourseRe
	for _, claim := range results.MyCourse {
		var ccin models.ClaimCourse
		var cc1 models.ClaimCourseRe
		var course models.CourseInsert
		fmt.Println("claim >>>> ", claim)
		err = s.DB(models.Database).C("claimcourse").Find(bson.M{"_id": claim}).One(&ccin)
		if err != nil {
			return user, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หา claimcourse ไม่ได้"}
		}
		err = s.DB(models.Database).C("course").Find(bson.M{"_id": ccin.IDCourse}).One(&course)
		if err != nil {
			course.ID = ccin.IDCourse
			err = nil
		}
		cc1.ID = ccin.ID
		cc1.IDUser = ccin.IDUser
		cc1.TimeEnd = ccin.TimeEnd
		t1 := time.Now()
		days := ccin.TimeEnd.Sub(t1).Hours() / 24
		cc1.TimeLeft = int(days)
		cc1.TimeStart = ccin.TimeStart
		cc1.Course = course
		cc = append(cc, cc1)
	}
	var order []models.Order
	for _, o := range results.Order {
		var oin models.Order
		err = s.DB(models.Database).C("order").Find(bson.M{"_id": o}).One(&oin)
		if err != nil {
			err = nil
		}
		order = append(order, oin)
	}

	user.MyCourse = cc
	user.Order = order

	return user, err
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

func UpdateMember(id string, u *models.RegisterRequest) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := models.MongoSession.Copy()
	defer s.Close()

	fmt.Println(u)
	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"email": u.Email,
		"firstname": u.FirstName, "lastname": u.LastName, "nickname": u.NickName,
		"telephonenumber": u.TelephoneNumber, "address": u.Address}}
	err := s.DB(models.Database).C("users").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
