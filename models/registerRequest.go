package models

type RegisterRequest struct {
	Username        string `json:"username" form:"username" query:"username"`
	Email           string `json:"email" form:"email" query:"email"`
	Password        string `json:"password" form:"password" query:"password"`
	FirstName       string `json:"firsname" form:"firsname" query:"firsname"`
	LastName        string `json:"lastname" form:"lastname" query:"lastname"`
	TelephoneNumber string `json:"telephonenumber" form:"telephonenumber" query:"telephonenumber"`
	NickName        string `json:"nickname" form:"nickname" query:"nickname"`
	Address         string `json:"address" form:"address" query:"address"`
	Bday            int    `json:"bday" form:"bday" query:"bday"`
	Bmonth          int    `json:"bmonth" form:"bmonth" query:"bmonth"`
	Byear           int    `json:"byear" form:"byear" query:"byear"`
}
