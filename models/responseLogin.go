package models

// ResponseLogin contains the token returned by the login
type ResponseLogin struct {
	Token string `json:"token,omitempty"`
}
