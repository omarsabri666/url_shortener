package user

type User struct {
	Email     string  `json:"email" validate:"required,email"`
	Password  string  `json:"password" validate:"required,min=7"`
	FirstName string  `json:"first_name" validate:"required,min=3"`
	LastName  string  `json:"last_name" validate:"required,min=3"`
	Id        *string `json:"id"`
}
type UserSignin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=7"`
}

type UserToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
type UserResponse struct {
	Data map[string]string `json:"data ,omitempty"`

	Message string `json:"message"`
	Success bool   `json:"success"`
}
