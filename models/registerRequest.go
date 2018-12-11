package models

type RegisterRequest struct {
	Username  string `json:"username" form:"username" query:"username"`
	Email     string `json:"email" form:"email" query:"email"`
	Password  string `json:"password" form:"password" query:"password"`
	FirstName string `json:"firsname" form:"firsname" query:"firsname"`
	LastName  string `json:"lastname" form:"lastname" query:"lastname"`
}
