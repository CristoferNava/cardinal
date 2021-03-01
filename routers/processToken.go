package routers

import (
	"errors"
	"strings"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email is the email of the user used in all the endpoints
var Email string

// UserID is the ID of the user used in all the endpoints
var UserID string

// ProcessToken process the token to extracts its values
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	key := []byte("Oliver_es_gay")
	claims := &models.Claim{}

	splitToken := strings.Split(token, "Bearer") // Bearer is for the JWT standart
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("token invalid format")
	}

	token = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err == nil {
		_, found, _ := db.CheckUserExists(claims.Email)
		if found {
			Email = claims.Email
			UserID = claims.ID.Hex()
		}
		return claims, found, UserID, nil
	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	return claims, false, string(""), err
}
