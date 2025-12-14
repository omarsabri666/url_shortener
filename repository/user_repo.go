	package repository

	import "github.com/omarsabri666/url_shorter/model/user"

	type UserRepository interface {
		CreateUser(user user.User) error
		GetUser(user user.UserSignin) (*user.User,error)
	}