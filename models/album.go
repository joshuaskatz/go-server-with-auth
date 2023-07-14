package models

type Album struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price int64 `json:"price"`
	CreatedAt string `json:"createdAt"`	
	UpdatedAt string `json:"updatedAt"`
}

