package models

type PremiumPackage struct {
	ID    int     `json:"id" db:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Quota int16   `json:"quota"`
}
