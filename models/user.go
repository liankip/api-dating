package models

import "time"

type User struct {
	ID               int       `json:"id" db:"id"`
	Username         string    `json:"username"`
	PasswordHash     string    `json:"-"`
	IsPremium        bool      `json:"is_premium"`
	VerifiedBadge    bool      `json:"verified_badge"`
	SwipeQuota       int       `json:"swipe_quota"`
	QuotaResetTime   time.Time `json:"quota_reset_time"`
	PremiumPackageID *int      `json:"premium_package_id"`
}
