package models

import "time"

type Swipe struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	ProfileID uint      `json:"profile_id"`
	Direction string    `json:"direction"`
	CreatedAt time.Time `json:"created_at"`
}
