package service

import (
	"github.com/fooku/LearnOnline_Api/models"
)

func AddComment(comment *models.Comment, id string) error {
	// if !bson.IsObjectIdHex(id) {
	// 	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	// }
	// idLec := bson.NewObjectId()
	// objectID := bson.ObjectIdHex(id)
	// s := models.MongoSession.Copy()
	// defer s.Close()

	// c := s.DB(models.Database).C("lectures")

	// lecturse.ID = idLec

	// var si models.SectionInsert
	// err := s.DB(models.Database).C("section").Find(bson.M{"_id": objectID}).One(&si)
	// if err != nil {
	// 	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	// }

	// err = c.Insert(&lecturse)
	// if err != nil {
	// 	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	// }

	// si.Lectures = append(si.Lectures, idLec)
	// colQuerier := bson.M{"_id": objectID}
	// change := bson.M{"$set": bson.M{"lectures": si.Lectures}}
	// err = s.DB(models.Database).C("section").Update(colQuerier, change)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	return nil
}
