package user

type User struct {
	Email     string  `json:"email" binding:"required,email"`
	Password  string  `json:"password" binding:"required,min=7"`
	FirstName string  `json:"first_name" binding:"required,min=3"`
	LastName  string  `json:"last_name" binding:"required,min=3"`
	Id        *string `json:"id"`
}
type UserSignin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=7"`
}

type UserToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
