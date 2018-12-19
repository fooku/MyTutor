package models

import (
	"gopkg.in/mgo.v2/bson"
)

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
}

func (cf *HomeContentFirst) AddContent() error {
	s := MongoSession.Copy()
	defer s.Close()

	c := s.DB(Database).C("homecontentfirst")

	err := c.Insert(&cf)

	return err
}

func (cs *HomeContentSecond) AddContent() error {
	s := MongoSession.Copy()
	defer s.Close()

	c := s.DB(Database).C("homecontentsecond")

	err := c.Insert(&cs)
	return err
}

func (ct *HomeContentThird) AddContent() error {
	s := MongoSession.Copy()
	defer s.Close()

	c := s.DB(Database).C("homecontentthird")

	err := c.Insert(&ct)
	return err
}
