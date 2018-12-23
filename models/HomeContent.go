package models

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

type HomeContent struct {
	HomeContentFirst  []HomeContentFirst  `json:"homecontentfirst"`
	HomeContentSecond []HomeContentSecond `json:"homecontentsecond"`
	HomeContentThird  []HomeContentThird  `json:"homecontentthird"`
}

// Collection: HomeContentFirst
// contentnumber ระบุลำดับการแสดงผลของเนื้อหา
// title ระบุหัวข้อของContent
// detail ระบุรายละเอียดของContent
// thumbnail ระบุ link ของรูปภาพของContent

// HomeContentFirst >> slind img content
type HomeContentFirst struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	ContentNumber uint8         `bson:"contentnumber" `
	Title         string        `bson:"title" `
	Detail        string        `bson:"detail" `
	Thumbnail     string        `bson:"thumbnail" `
}

// Collection: HomeContentSecond
// contentnumber ระบุลำดับการแสดงผลของเนื้อหา
// title ระบุหัวข้อของContent
// detail ระบุรายละเอียดของContent
// icon ระบุ icon ของรูปภาพของContent

type HomeContentSecond struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	ContentNumber uint8         `bson:"contentnumber" `
	Title         string        `bson:"title" `
	Detail        string        `bson:"detail" `
	Icon          string        `bson:"icon" `
}

// Collection: HomeContentThird
// contentnumber ระบุลำดับการแสดงผลของเนื้อหา
// title ระบุหัวข้อของContent
// detail ระบุรายละเอียดของContent
// thumbnail ระบุ link ของรูปภาพของContent

type HomeContentThird struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	ContentNumber uint8         `bson:"contentnumber"`
	Title         string        `bson:"title" `
	Detail        string        `bson:"detail" `
	Thumbnail     string        `bson:"thumbnail" `
}

type AddHomeContent interface {
	AddContent() error
	DeleteContent(string) error
}
type UpdateHomeContent interface {
	UpdateContent(string) error
}

func (cf *HomeContentFirst) AddContent() error {
	s := MongoSession.Copy()
	defer s.Close()

	c := s.DB(Database).C("homecontentfirst")

	err := c.Insert(&cf)

	return err
}

func (ct *HomeContentThird) AddContent() error {
	s := MongoSession.Copy()
	defer s.Close()

	c := s.DB(Database).C("homecontentthird")

	err := c.Insert(&ct)
	return err
}

func (cf *HomeContentFirst) UpdateContent(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"contentnumber": cf.ContentNumber, "title": cf.Title, "detail": cf.Detail, "thumbnail": cf.Thumbnail}}
	err := s.DB(Database).C("homecontentfirst").Update(colQuerier, change)
	return err
}
func (cs *HomeContentSecond) UpdateContent(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	fmt.Println("id >><><><>", id)
	objectID := bson.ObjectIdHex(id)
	s := MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"contentnumber": cs.ContentNumber, "title": cs.Title, "detail": cs.Detail, "icon": cs.Icon}}
	err := s.DB(Database).C("homecontentsecond").Update(colQuerier, change)
	return err
}
func (ct *HomeContentThird) UpdateContent(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"contentnumber": ct.ContentNumber, "title": ct.Title, "detail": ct.Detail, "thumbnail": ct.Thumbnail}}
	err := s.DB(Database).C("homecontentthird").Update(colQuerier, change)
	return err
}

func (cf *HomeContentFirst) DeleteContent(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := MongoSession.Copy()
	defer s.Close()

	err := s.DB(Database).C("homecontentfirst").RemoveId(objectID)
	return err
}
func (ct *HomeContentThird) DeleteContent(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(id)
	s := MongoSession.Copy()
	defer s.Close()

	err := s.DB(Database).C("homecontentthird").RemoveId(objectID)
	return err
}
