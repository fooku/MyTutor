package models

import (
	"gopkg.in/mgo.v2"
)

var MongoSession *mgo.Session

const Database = "gutututor"

// Init mongodb
func Init(mongoURL string) error {
	var err error
	MongoSession, err = mgo.Dial(mongoURL)
	if err != nil {
		return err
	}
	return nil
}
