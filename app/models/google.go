package models

import "time"

type GoogleUrlResponse struct {
	Url string `json:"url"`
}

type GoogleCodeRequest struct {
	Code string `json:"code"`
}

type GoogleProfile struct {
	Id            string    `json:"id"`
	Email         string    `json:"email"`
	VerifiedEmail bool      `json:"verified_email"`
	Name          string    `json:"name"`
	Picture       string    `json:"picture"`
	AccessToken   string    `json:"access_token"`
	TokenType     string    `json:"token_type,omitempty"`
	RefreshToken  string    `json:"refresh_token,omitempty"`
	Expiry        time.Time `json:"expiry,omitempty"`
}
