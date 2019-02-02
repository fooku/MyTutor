package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type ClaimCourse struct {
	ID        bson.ObjectId `json:"id" bson:"_id,omitempty"`
	IDUser    bson.ObjectId `json:"iduser" bson:"iduser,omitempty"`
	IDCourse  bson.ObjectId `json:"idcourse" bson:"idcourse,omitempty" `
	TimeStart time.Time
	TimeEnd   time.Time
}
