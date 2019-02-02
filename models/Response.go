package models

type LoginResponse struct {
	User  UserResponse `json:"user" form:"user" query:"user"`
	Token string       `json:"token" form:"token" query:"token"`
}
