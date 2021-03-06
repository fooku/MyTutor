package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name" `
	Hour      string        `json:"hour" bson:"hour" `
	Price     string        `json:"price" bson:"price" `
	Publish   bool          `json:"publish" bson:"publish" `
	ClaimUser []ClaimCourse `json:"claimuser" bson:"claimuser" `
	Type      string        `json:"type" bson:"type" `
	Detail    string        `json:"detail" bson:"detail" `
	Thumbnail string        `json:"thumbnail" bson:"thumbnail" `
	Section   []Section     `json:"section" bson:"section" `
	Timestamp time.Time
}

type CourseInsert struct {
	ID        bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name      string          `json:"name" bson:"name" `
	Hour      string          `json:"hour" bson:"hour" `
	Creator   bson.ObjectId   `json:"creator" son:"creator" `
	Price     string          `json:"price" bson:"price" `
	Publish   bool            `json:"publish" bson:"publish" `
	ClaimUser []bson.ObjectId `json:"claimuser" bson:"claimuser" `
	Type      string          `json:"type" bson:"type" `
	Detail    string          `json:"detail" bson:"detail" `
	Thumbnail string          `json:"thumbnail" bson:"thumbnail" `
	Section   []bson.ObjectId `json:"section" bson:"section" `
	Timestamp time.Time
}
