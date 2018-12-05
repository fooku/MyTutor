package api

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/fooku/authBasic/models"
	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
)

// Accessible > หน้าแรก
func Accessible(c echo.Context) error {
	return c.String(http.StatusOK, "Accessible")
}

// Login > ยืนยันตัวต้นจาก HasPassword และ เจน jwt key
func Login(c echo.Context) error {
	u := new(models.LoginRequest)
	err := c.Bind(u)

	err, user := models.FindUser(u.Username)

	if err == mgo.ErrNotFound {
		return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid email"}
	}

	if user.ComparePassword(u.Password) {
		// Create token
		token := jwt.New(jwt.SigningMethodHS256)

		// Set claims
		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		// Generate encoded token and send it as response.
		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return &echo.HTTPError{Code: http.StatusUnauthorized, Message: "invalid password"}
}

// Register > เพิ่มผู้ใช้ลง mongoDB
func Register(c echo.Context) error {
	var user models.User

	u := new(models.RegisterRequest)
	err := c.Bind(u)

	fmt.Println(u)
	err = models.AddUser(user, u.Username, u.Email, u.Password)

	return err
}

// Restricted > jwt
func Restricted(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)

	// เพิ่มคำร้อง

	fmt.Println(name, claims["admin"])
	return c.JSON(http.StatusOK, claims)
}
