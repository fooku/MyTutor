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

func GetHomeContent() (*models.HomeContent, error) {
	content := new(models.HomeContent)
	s := models.MongoSession.Copy()
	defer s.Close()

	var first []models.HomeContentFirst
	var second []models.HomeContentSecond
	var third []models.HomeContentThird
	err := s.DB(models.Database).C("homecontentfirst").Find(nil).All(&first)
	if err != nil {
		return content, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	err = s.DB(models.Database).C("homecontentfirst").Find(nil).All(&second)
	if err != nil {
		return content, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	err = s.DB(models.Database).C("homecontentfirst").Find(nil).All(&third)
	if err != nil {
		return content, &echo.HTTPError{Code: http.StatusUnauthorized, Message: err}
	}
	content.HomeContentFirst = first
	content.HomeContentSecond = second
	content.HomeContentThird = third

	fmt.Println("Results HomeContent: ", content)
	return content, err
}
