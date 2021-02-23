package jwt

import (
	"time"

	"github.com/CristoferNava/cardinal/models"
	jwt "github.com/dgrijalva/jwt-go"
)

// GenerateJWT generates a token for a user based on our private key
func GenerateJWT(user models.User) (string, error) {
	key := []byte("Oliver_es_gay")

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"lastname":  user.LastName,
		"birthdate": user.Birthdate,
		"bio":       user.Bio,
		"location":  user.Location,
		"web":       user.Web,
		"_id":       user.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(key)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, nil
}
