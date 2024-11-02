package infrastructure

import (
	"api-dating/models"
	"database/sql"
	"fmt"
	"log"
)

func SeedPremiumPackages(db *sql.DB) error {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM seed_history WHERE name = 'SeedPremiumPackages')`
	err := db.QueryRow(query).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking seed history: %w", err)
	}

	if exists {
		log.Println("SeedPremiumPackages has already been run, skipping.")
		return nil
	}

	packages := []models.PremiumPackage{
		{Name: "Basic", Price: 50.000, Quota: 200},
		{Name: "Premium", Price: 100.000, Quota: 300},
		{Name: "Elite", Price: 150.000, Quota: 500},
	}

	for _, pkg := range packages {
		_, err := db.Exec(`
			INSERT INTO premium_packages (name, price, quota)
			VALUES ($1, $2, $3)`, pkg.Name, pkg.Price, pkg.Quota)
		if err != nil {
			return fmt.Errorf("error seeding premium_packages: %w", err)
		}
	}

	_, err = db.Exec(`INSERT INTO seed_history (name) VALUES ('SeedPremiumPackages')`)
	if err != nil {
		return fmt.Errorf("error recording seed history: %w", err)
	}

	log.Println("SeedPremiumPackages executed successfully.")
	return nil
}
