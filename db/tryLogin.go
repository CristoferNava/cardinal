package db

import (
	"github.com/CristoferNava/cardinal/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin checks the credentials against the DB
func TryLogin(email, password string) (models.User, bool) {
	user, found, _ := CheckUserExists(email)
	if !found {
		return user, false
	}

	// check the passwords
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}
