package models

// JWT is the struct to send the token to the client
type JWT struct {
	Token string `json:"token,omitempty"`
}
