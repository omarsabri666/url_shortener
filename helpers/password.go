package helpers

import (
	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string , salt int )(string,error){
	hashedPassword,err := bcrypt.GenerateFromPassword([]byte(password),salt)
	if err != nil {
		return "",err
	}
	return string(hashedPassword),nil
}
func ComparePassword(passwordFromReq string , passwordFromDb string ) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(passwordFromDb),[]byte(passwordFromReq))
	if err != nil {
		return err
	}
	return nil

}