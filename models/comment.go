package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Comment struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Sender    bson.ObjectId `json:"sender" bson:"sender,omitempty"`
	Content   string        `json:"content" bson:"content" `
	Timestamp time.Time
}
