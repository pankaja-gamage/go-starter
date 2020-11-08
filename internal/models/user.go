package models

//User structure
type User struct {
	Model
	Name   string
	Email  string
	Mobile string
}

//AddUserRequest structure
type AddUserRequest struct {
	Name   string
	Email  string
	Mobile string
}
