package model

type User struct {
	Id int `json:"id"`
	Name    string `json:"name" orm:"size(128)"`
	Email    string `json:"email" orm:"size(128)"`
	Password string `json:"password" orm:"size(64)"`
	Username string `json:"username" orm:"size(32)"`
}
