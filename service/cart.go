package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AddCart(iduser string, item *models.Item) error {
	fmt.Println("id", iduser)
	if !bson.IsObjectIdHex(iduser) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(iduser)
	s := models.MongoSession.Copy()
	defer s.Close()

	var result models.User
	err := s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&result)
	if err != nil {
		return err
	}

	resultFind := find(result.Cart, item.Course_ID)
	if !resultFind {
		result.Cart = append(result.Cart, *item)
	}

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"cart": result.Cart}}
	err = s.DB(models.Database).C("users").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func DeleteItem(iduser string, i int) error {
	fmt.Println("id", iduser)
	if !bson.IsObjectIdHex(iduser) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(iduser)
	s := models.MongoSession.Copy()
	defer s.Close()

	var result models.User
	err := s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&result)
	if err != nil {
		return err
	}

	result.Cart = append(result.Cart[:i], result.Cart[i+1:]...)

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"cart": result.Cart}}
	err = s.DB(models.Database).C("users").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func GenOrder(iduser string) error {
	fmt.Println("id", iduser)
	if !bson.IsObjectIdHex(iduser) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(iduser)
	s := models.MongoSession.Copy()
	defer s.Close()

	var result models.User
	err := s.DB(models.Database).C("users").Find(bson.M{"_id": objectID}).One(&result)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Find user  ไม่ได้"}
	}

	sum := 0
	for _, item := range result.Cart {
		sum += item.Price
	}

	var order models.Order
	order.ID = bson.NewObjectId()
	order.Cart = result.Cart
	order.Total = sum
	order.UserID = iduser
	order.Status = "รอโอนเงิน"
	order.Timestamp = time.Now()
	err = s.DB(models.Database).C("order").Insert(&order)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Insert order ไม่ได้"}
	}

	result.Order = append(result.Order, order.ID)

	colQuerier := bson.M{"_id": objectID}
	var item []models.Item
	change := bson.M{"$set": bson.M{"cart": item, "order": result.Order}}
	err = s.DB(models.Database).C("users").Update(colQuerier, change)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update cart in users ไม่ได้"}
	}

	return nil
}

func find(items []models.Item, key string) bool {
	for _, item := range items {
		if item.Course_ID == key {
			return true
		}
	}
	return false
}
