package models

type AuthResponse struct {
	UserId	   uint `json:"user_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
