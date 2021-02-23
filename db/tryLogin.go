package db

import (
	"github.com/CristoferNava/cardinal/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin against the database
func TryLogin(email, password string) (models.User, bool) {
	user, found, _ := CheckUserExists(email)
	if found == false {
		return user, false
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
