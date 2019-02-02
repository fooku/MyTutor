package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	HasPassword
	UserType        string          `bson:"usertype"`
	Email           string          `bson:"email"`
	FirstName       string          `bson:"firstname"`
	LastName        string          `bson:"lastname"`
	NickName        string          `bson:"nickname"`
	TelephoneNumber string          `bson:"telephonenumber"`
	Address         string          `bson:"address"`
	MyCourse        []bson.ObjectId `bson:"mycourse"`
	Birthday        time.Time
	Timestamp       time.Time
}

type UserResponse struct {
	ID bson.ObjectId `bson:"_id,omitempty"`
	HasPassword
	UserType        string        `bson:"usertype"`
	Email           string        `bson:"email"`
	FirstName       string        `bson:"firstname"`
	LastName        string        `bson:"lastname"`
	NickName        string        `bson:"nickname"`
	TelephoneNumber string        `bson:"telephonenumber"`
	Address         string        `bson:"address"`
	MyCourse        []ClaimCourse `bson:"mycourse"`
	Birthday        time.Time
	Timestamp       time.Time
}
