package models

type RegisterRequest struct {
	Email           string `json:"email" form:"email" query:"email"`
	Password        string `json:"password" form:"password" query:"password"`
	FirstName       string `json:"firstname" form:"firstname" query:"firstname"`
	LastName        string `json:"lastname" form:"lastname" query:"lastname"`
	TelephoneNumber string `json:"telephonenumber" form:"telephonenumber" query:"telephonenumber"`
	NickName        string `json:"nickname" form:"nickname" query:"nickname"`
	Address         string `json:"address" form:"address" query:"address"`
	Bday            int    `json:"bday" form:"bday" query:"bday"`
	Bmonth          int    `json:"bmonth" form:"bmonth" query:"bmonth"`
	Byear           int    `json:"byear" form:"byear" query:"byear"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}

type UpdateRequest struct {
	Usertype string `json:"usertype" form:"usertype" query:"usertype"`
}

type ResetRequest struct {
	Email       string `json:"email" form:"email" query:"email"`
	Password    string `json:"password" form:"password" query:"password"`
	NewPassword string `json:"newpassword" form:"newpassword" query:"newpassword"`
}
