package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts the password using bcrypt
func EncryptPassword(password string) (string, error) {
	cost := 6 // number of iterations 2^cost
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	return string(bytes), err
}
