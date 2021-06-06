package jwt

import (
	"time"

	"github.com/CristoferNava/cardinal/models"
	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(user models.User) (string, error) {
	privateKey := []byte("laMayonesaNoEsUnInstrumento")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastName":  user.LastName,
		"birthdate": user.Birthdate,
		"bio":       user.Bio,
		"location":  user.Location,
		"website":   user.Website,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
