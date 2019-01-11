package models

import "gopkg.in/mgo.v2/bson"

type Lectures struct {
	ID      bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name    string        `json:"name" bson:"name" `
	Time    float64       `json:"time" bson:"time" `
	Comment string        `json:"comment" bson:"comment" `
}
