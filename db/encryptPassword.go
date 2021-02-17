package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword encrypts the user's password
func EncryptPassword(password string) (string, error) {
	cost := 8 // 2^cost times to encrypt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
