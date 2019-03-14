package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Order struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Cart      []Item        `bson:"cart" json:"cart"`
	UserID    string        `bson:"userid" json:"userid"`
	Total     int           `bson:"total" json:"total"`
	Status    string        `bson:"status" json:"status"`
	Payment   Payment       `bson:"payment" json:"payment"`
	Payment2  string        `bson:"payment2" json:"payment2"`
	Timestamp time.Time     `json:"timestamp"`
}

type OrderRe struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Cart      []Item        `bson:"cart" json:"cart"`
	User      User          `bson:"userid" json:"userid"`
	Total     int           `bson:"total" json:"total"`
	Status    string        `bson:"status" json:"status"`
	Payment   Payment       `bson:"payment" json:"payment"`
	Payment2  string        `bson:"payment2" json:"payment2"`
	Timestamp time.Time     `json:"timestamp"`
}

type Payment struct {
	TelephoneNumber string    `bson:"telephonenumber" json:"telephonenumber"`
	Bank            string    `bson:"bank" json:"bank"`
	Date            time.Time `bson:"date" json:"date"`
	Time            string    `bson:"time" json:"time"`
	Total           int       `bson:"total" json:"total"`
	image           string    `bson:"image" json:"image"`
}
