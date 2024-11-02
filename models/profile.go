package models

type Profile struct {
	ID       int    `json:"id" db:"id"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Bio      string `json:"bio"`
	PhotoURL string `json:"photo_url"`
}
