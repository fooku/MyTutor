package service

import (
	"fmt"
	"net/http"

	"github.com/fooku/LearnOnline_Api/models"
	"github.com/labstack/echo"
)

func AddHomeContent(ahc models.AddHomeContent) error {
	err := ahc.AddContent()
	if err != nil {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	return nil
}

func GetHomeContentFirst() ([]models.HomeContentFirst, error) {
	s := models.MongoSession.Copy()
	defer s.Close()

	var results []models.HomeContentFirst
	err := s.DB(models.Database).C("homecontentfirst").Find(nil).All(&results)

	fmt.Println("Results All: ", results)
	return results, err
}
