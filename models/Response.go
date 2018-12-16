package models

type LoginResponse struct {
	User  User   `json:"user" form:"user" query:"user"`
	Token string `json:"token" form:"token" query:"token"`
}
