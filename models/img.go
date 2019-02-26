package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Img struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Img       string        `json:"img" bson:"img" `
	Timestamp time.Time
}
