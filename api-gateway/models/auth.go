package models

type Creds struct {
	Username string
	Password Password
}

type AuthResponse struct {
	Token string `json:"token"`
}
