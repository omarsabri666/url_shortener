package service

import (
	"testing"

	"github.com/omarsabri666/url_shorter/helpers"
)

// import (
// 	"testing"

// 	"github.com/omarsabri666/url_shorter/model/user"
// 	// "github.com/omarsabri666/url_shorter/repository"
// )

// func TestLogin_success(t *testing.T) {
// 	service := NewUserService(newFakeUserRepo(fakeUserRepo{user: &user.User{Email: "email",Password: "password"}, err: nil}),fakeTokenRepo)

// }
// type fakeUserRepo struct {
// 	user *user.User
// 	err error

// }
// func newFakeUserRepo( st fakeUserRepo) *fakeUserRepo {
// 	return &fakeUserRepo{user: st.user, err: st.err}
// }
// func (f * fakeUserRepo) GetUser(user user.UserSignin) (*user.User, error) {
// 	return  f.user,f.err

// }
// func (f * fakeUserRepo) CreateUser(user user.User) error {
// 	return f.err
// }
// 	type fakeTokenRepo struct {

// 	}
















func TestComparePassword_success(t *testing.T) {
	password := "password"
	passwordToCompare := "password"
hashedPassword, err :=	helpers.HashPassword(passwordToCompare,10)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

 err = 	helpers.ComparePassword(password,hashedPassword)
	if err != nil {
		t.Errorf("Error comparing password: %v", err)
	}
}
func TestComparePassword_fail(t *testing.T) {
	password := "password1"
	passwordToCompare := "password"
hashedPassword, err :=	helpers.HashPassword(passwordToCompare,10)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}

 err = 	helpers.ComparePassword(password,hashedPassword)
	if err == nil {
		t.Errorf("Error comparing password: %v", err)
	}
}