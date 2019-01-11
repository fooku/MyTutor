package models

import "gopkg.in/mgo.v2/bson"

type Section struct {
	ID       bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name     string        `json:"name" bson:"name" `
	Lectures []Lectures    `json:"lectures" bson:"lectures" `
}

type SectionInsert struct {
	ID       bson.ObjectId   `json:"id" bson:"_id,omitempty"`
	Name     string          `json:"name" bson:"name" `
	Lectures []bson.ObjectId `json:"lectures,omitempty" bson:"lectures,omitempty" `
}
