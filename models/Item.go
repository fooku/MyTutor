package models

type Item struct {
	Course_ID string `json:"courseid" bson:"course" `
	Name      string `json:"name" bson:"name" `
	Price     int    `json:"price" bson:"price" `
	Img       string `json:"img" bson:"img" `
}
