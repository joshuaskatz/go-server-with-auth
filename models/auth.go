package models

type AuthInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Verified bool   `json:"verified"`
}

type EmailTemplate struct {
	Email string
}
