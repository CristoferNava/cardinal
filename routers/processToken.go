package routers

import (
	"errors"
	"strings"

	"github.com/CristoferNava/cardinal/db"
	"github.com/CristoferNava/cardinal/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// Email is the email of the user
var Email string

// IDUser is the _id of the user
var IDUser string

// ProcessToken given a token validates if is correct or not
func ProcessToken(tokenToValidate string) (*models.Claim, bool, string, error) {
	privateKey := []byte("laMayonesNoEsUnInstrumento")
	claims := &models.Claim{}

	splittedToken := strings.Split(tokenToValidate, "Bearer")
	if len(splittedToken) != 2 {
		return claims, false, "", errors.New("Invalid Token format")
	}

	tokenToValidate = strings.TrimSpace(splittedToken[1])
	validatedToken, err := jwt.ParseWithClaims(tokenToValidate, claims, func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	}) // validate the token and save the data in the claims struct
	if err != nil {
		if !validatedToken.Valid {
			return claims, false, "", errors.New("Invalid Token")
		}
		return claims, false, "", err
	}

	_, found, _ := db.CheckUserExists(claims.Email)
	if found {
		Email = claims.Email
		IDUser = claims.ID.Hex()
	}
	return claims, found, IDUser, nil
}
