package password

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func CreateHashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}

func VerifyPassword(userPassWord string, givenPassWord string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(givenPassWord), []byte(userPassWord)) 
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}