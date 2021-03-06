package api

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/LearnOnline_Api/models"
	"github.com/fooku/LearnOnline_Api/service"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Accessible > หน้าแรก
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Login > ยืนยันตัวต้นจาก HasPassword และ เจน jwt key
func Login(c echo.Context) error {
	u := new(models.LoginRequest)
	err := c.Bind(u)

	err, user := service.FindUser(u.Email)

	if err == mgo.ErrNotFound {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email"}
	}

	if user.ComparePassword(u.Password) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["UserType"] = user.UserType
		claims["exp"] = time.Now().Add(time.Hour * 7200).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		res := new(models.LoginResponse)

		res.User = user
		res.Token = t
		return c.JSON(http.StatusOK, res)
	}

	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid password"}
}

// Register > เพิ่มผู้ใช้ลง mongoDB
func Register(c echo.Context) error {
	var user models.User

	u := new(models.RegisterRequest)
	err := c.Bind(u)

	fmt.Println(u)
	//เก็บข้อมูลผู้ใช้ลง mongoDB
	err = service.AddUser(&user, u)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"Message": "Succeed",
	})
}

func ResetPassword2(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	if claims["UserType"] != "admin" {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "admin only naja"}
	}

	u := new(models.ResetRequest)
	err := c.Bind(u)

	err, userf := service.FindUser(u.Email)

	if err == mgo.ErrNotFound {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email"}
	}

	NewPassword := RandStringRunes(10)
	userf.SetPassword(NewPassword)

	err = service.ResetPassword(&userf)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]string{
		"newpassword": NewPassword,
	})
}

// Restricted > jwt
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	if !bson.IsObjectIdHex(id) {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid id"}
	}
	u, err := service.GetMemberOne(id)
	if err != nil {
		return err
	}
	// เพิ่มคำร้อง
	return c.JSON(http.StatusOK, u)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
