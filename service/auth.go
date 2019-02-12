package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// AddUser > insert data to mongoDB
func AddUser(user *models.User, u *models.RegisterRequest) error {
	s := models.MongoSession.Copy()
	defer s.Close()

	c := s.DB(models.Database).C("users")

	index := mgo.Index{
		Key:    []string{"telephonenumber"},
		Unique: true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		fmt.Println(err)
	}

	user.SetPassword(u.Password)
	user.Email = u.Email
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.UserType = "member"
	user.NickName = u.NickName
	user.TelephoneNumber = u.TelephoneNumber
	user.Address = u.Address
	user.Birthday = time.Date(
		u.Byear, time.Month(u.Bmonth), u.Bday, 0, 0, 0, 0, time.UTC)
	user.Timestamp = time.Now()
	err = c.Insert(&user)
	// err := s.DB(database).C("users").Update("Username", &user)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func FindUser(email string) (error, models.UserResponse) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var user models.UserResponse

	var result models.User
	err := s.DB(models.Database).C("users").Find(bson.M{"email": email}).One(&result)
	// Find(bson.M{"username": "sdfasdfasdf"}).Sort("-timestamp").All(&results)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาUserone ไม่ได้"}, user
	}

	var cc []models.ClaimCourse
	for _, claim := range result.MyCourse {
		var cc1 models.ClaimCourse
		fmt.Println("claim >>>> ", claim)
		err = s.DB(models.Database).C("claimcourse").Find(bson.M{"_id": claim}).One(&cc1)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หา claimcourse ไม่ได้"}, user
		}
		cc = append(cc, cc1)
	}

	user.ID = result.ID
	user.UserType = result.UserType
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.NickName = result.NickName
	user.TelephoneNumber = result.TelephoneNumber
	user.Timestamp = result.Timestamp
	user.Address = result.Address
	user.Birthday = result.Birthday
	user.HasPassword = result.HasPassword
	user.MyCourse = cc

	fmt.Println("Results All: ", result.HasPassword)

	return nil, user
}
