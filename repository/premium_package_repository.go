// repository/premium_repo.go
package repository

import (
	"api-dating/models"
	"database/sql"
	"errors"
)

type PremiumPackageRepository interface {
	GetPremiumPackageByID(id int) (*models.PremiumPackage, error)

	ActivateUserPremium(userID uint, packageID int) error
}

type PremiumPackageRepositoryImpl struct {
	db *sql.DB
}

func NewPremiumPackageRepository(db *sql.DB) PremiumPackageRepository {
	return &PremiumPackageRepositoryImpl{db: db}
}

func (repository *PremiumPackageRepositoryImpl) GetPremiumPackageByID(id int) (*models.PremiumPackage, error) {
	packageData := &models.PremiumPackage{}

	query := `SELECT id, name, price, quota FROM premium_packages WHERE id = $1`

	err := repository.db.QueryRow(query, id).Scan(&packageData.ID, &packageData.Name, &packageData.Price, &packageData.Quota)

	if err == sql.ErrNoRows {
		return nil, errors.New("premium package not found")
	}

	return packageData, err
}

func (repository *PremiumPackageRepositoryImpl) ActivateUserPremium(userID uint, packageID int) error {
	query := `UPDATE users SET is_premium = $1, premium_package_id = $2 WHERE id = $3`

	_, err := repository.db.Exec(query, true, packageID, userID)

	return err
}
