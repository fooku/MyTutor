package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type News struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Title     string        `bson:"title" `
	Detail    string        `bson:"detail" `
	Thumbnail string        `bson:"thumbnail" `
	Timestamp time.Time
}
