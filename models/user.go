package models


type User struct {
    ID string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role     string `json:"role"`
    CreatedAt string `json:"createdAt"`	
	UpdatedAt string `json:"updatedAt"`
}