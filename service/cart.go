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

func AddPayment(idoder string, payment *models.Payment) error {
	fmt.Println("id", idoder)
	if !bson.IsObjectIdHex(idoder) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(idoder)
	s := models.MongoSession.Copy()
	defer s.Close()

	// var result models.Order
	// err := s.DB(models.Database).C("order").Find(bson.M{"_id": objectID}).One(&result)
	// if err != nil {
	// 	return err
	// }

	// result.Payment = payment

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"payment": payment, "status": "รอเช็คยอดเงิน"}}
	err := s.DB(models.Database).C("order").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
func GetOrderOne(idOrder string) (*models.Order, error) {
	s := models.MongoSession.Copy()
	defer s.Close()
	if !bson.IsObjectIdHex(idOrder) {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	id := bson.ObjectIdHex(idOrder)
	var order models.Order

	err := s.DB(models.Database).C("order").Find(bson.M{"_id": id}).One(&order)
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาorder one ไม่เจอ"}
	}

	return &order, nil
}

func ListOrder() (*[]models.Order, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var order []models.Order

	err := s.DB(models.Database).C("order").Find(nil).All(&order)
	if err != nil {
		return nil, &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาorder ไม่เจอ"}
	}

	return &order, nil
}

func AddPayment2(idoder string, img string) error {
	fmt.Println("id", idoder)
	if !bson.IsObjectIdHex(idoder) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	objectID := bson.ObjectIdHex(idoder)
	s := models.MongoSession.Copy()
	defer s.Close()

	colQuerier := bson.M{"_id": objectID}
	change := bson.M{"$set": bson.M{"payment2": img, "status": "รอเช็คยอดเงิน"}}
	err := s.DB(models.Database).C("order").Update(colQuerier, change)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func BuyMany(idOrder string, date int) error {
	s := models.MongoSession.Copy()
	defer s.Close()

	id := bson.ObjectIdHex(idOrder)
	var order models.Order

	err := s.DB(models.Database).C("order").Find(bson.M{"_id": id}).One(&order)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาorder one ไม่เจอ"}
	}

	objectIDUser := bson.ObjectIdHex(order.UserID)

	var user models.User
	colQuerierUser := bson.M{"_id": objectIDUser}
	err = s.DB(models.Database).C("users").Find(bson.M{"_id": objectIDUser}).One(&user)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาUser one ไม่ได้"}
	}

	for _, item := range order.Cart {
		objectIDCourse := bson.ObjectIdHex(item.Course_ID)
		var cci models.CourseInsert
		colQuerier := bson.M{"_id": objectIDCourse}
		err = s.DB(models.Database).C("course").Find(bson.M{"_id": objectIDCourse}).One(&cci)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "หาCourse one ไม่ได้"}
		}

		var cu models.ClaimCourse
		cu.ID = bson.NewObjectId()
		cu.IDCourse = objectIDCourse
		cu.IDUser = objectIDUser
		cu.TimeStart = time.Now()
		cu.TimeEnd = time.Now().AddDate(0, 0, date)
		err = s.DB(models.Database).C("claimcourse").Insert(&cu)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Insert claimcourse ไม่ได้"}
		}

		cci.ClaimUser = append(cci.ClaimUser, cu.ID)
		change := bson.M{"$set": bson.M{"claimuser": cci.ClaimUser}}
		err = s.DB(models.Database).C("course").Update(colQuerier, change)
		if err != nil {
			fmt.Println(err)
		}

		user.MyCourse = append(user.MyCourse, cu.ID)
		changeUser := bson.M{"$set": bson.M{"mycourse": user.MyCourse}}
		err = s.DB(models.Database).C("users").Update(colQuerierUser, changeUser)
		if err != nil {
			return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update mycourse in course ไม่ได้"}
		}
	}

	colQuerier := bson.M{"_id": id}
	change := bson.M{"$set": bson.M{"status": "คำสั่งซื้อถูกอนุมัติแล้ว"}}
	err = s.DB(models.Database).C("order").Update(colQuerier, change)
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "Update order status ไม่ได้"}
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
