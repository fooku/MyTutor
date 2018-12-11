package models

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	HasPassword
	UserType  string `bson:"usertype"`
	Email     string `bson:"email"`
	Username  string `bson:"username"`
	FirstName string `bson:"firsname"`
	LastName  string `bson:"lastname"`
	Timestamp time.Time
}

// AddUser > insert data to mongoDB
func AddUser(user User, username string, email string, password string, firsname string, lastname string) error {
	s := mongoSession.Copy()
	defer s.Close()

	c := s.DB(database).C("users")

	index := mgo.Index{
		Key:    []string{"username", "email"},
		Unique: true,
	}

	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	user.SetPassword(password)
	user.Username = username
	user.Email = email
	user.FirstName = firsname
	user.LastName = lastname
	user.UserType = "member"
	user.Timestamp = time.Now()
	err = c.Insert(&user)
	// err := s.DB(database).C("users").Update("Username", &user)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func FindUser(email string) (error, User) {
	s := mongoSession.Copy()
	defer s.Close()

	var result User
	err := s.DB(database).C("users").Find(bson.M{"email": email}).One(&result)
	// Find(bson.M{"username": "sdfasdfasdf"}).Sort("-timestamp").All(&results)

	if err != nil {
		return err, result
	}

	fmt.Println("Results All: ", result.HasPassword)

	return nil, result
}

func GetMember() ([]User, error) {
	s := mongoSession.Copy()
	defer s.Close()

	var results []User
	err := s.DB(database).C("users").Find(nil).Sort("-timestamp").All(&results)

	fmt.Println("Results All: ", results)
	return results, err
}

func DeleteMember(id string) error {
	fmt.Println("id", id)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := mongoSession.Copy()
	defer s.Close()
	err := s.DB(database).C("users").RemoveId(objectID)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusNotFound, Message: "not found"}
	}
	return nil
}
