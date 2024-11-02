// repository/swipe_repo.go
package repository

import (
	"api-dating/models"
	"database/sql"
	"time"
)

type SwipeRepository interface {
	RecordSwipe(swipe *models.Swipe) error

	HasSwipedToday(userID uint, profileID int) (bool, error)
}

type SwipeRepoImpl struct {
	db *sql.DB
}

func NewSwipeRepository(db *sql.DB) SwipeRepository {
	return &SwipeRepoImpl{db: db}
}

func (repository *SwipeRepoImpl) RecordSwipe(swipe *models.Swipe) error {
	query := `INSERT INTO swipes (user_id, profile_id, direction, created_at)
			  VALUES ($1, $2, $3, $4)`

	_, err := repository.db.Exec(query, swipe.UserID, swipe.ProfileID, swipe.Direction, time.Now())

	return err
}

func (repository *SwipeRepoImpl) HasSwipedToday(userID uint, ProfileID int) (bool, error) {
	query := `SELECT COUNT(1) FROM swipes 
			  WHERE user_id = $1 AND profile_id = $2 AND DATE(created_at) = CURRENT_DATE`

	var count int

	err := repository.db.QueryRow(query, userID, ProfileID).Scan(&count)

	if err != nil {
		return false, err
	}

	return count > 0, nil
}
