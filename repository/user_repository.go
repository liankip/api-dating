// repository/user_repo.go
package repository

import (
	"api-dating/models"
	"database/sql"
	"errors"
)

type UserRepository interface {
	CreateUser(user *models.User) error

	FindUserByUsername(username string) (*models.User, error)

	UpdateSwipeQuota(userID uint, newQuota int) error

	GetUserByID(userID uint) (*models.User, error)
}

type UserRepoImpl struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &UserRepoImpl{db: db}
}

func (repository *UserRepoImpl) CreateUser(user *models.User) error {
	query := `INSERT INTO users (username, password_hash, is_premium, verified_badge, swipe_quota, quota_reset_time)
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	err := repository.db.QueryRow(query, user.Username, user.PasswordHash, user.IsPremium, user.VerifiedBadge, user.SwipeQuota, user.QuotaResetTime).Scan(&user.ID)

	return err
}

func (repository *UserRepoImpl) FindUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `SELECT id, username, password_hash, is_premium, verified_badge, swipe_quota, quota_reset_time 
			  FROM users WHERE username = $1`

	err := repository.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.PasswordHash, &user.IsPremium, &user.VerifiedBadge, &user.SwipeQuota, &user.QuotaResetTime)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (repository *UserRepoImpl) UpdateSwipeQuota(userID uint, newQuota int) error {
	query := `UPDATE users SET swipe_quota = $1 WHERE id = $2`

	_, err := repository.db.Exec(query, newQuota, userID)

	return err
}

func (repository *UserRepoImpl) GetUserByID(userID uint) (*models.User, error) {
	user := &models.User{}

	query := `SELECT id, username, password_hash, is_premium, verified_badge, swipe_quota, quota_reset_time 
			  FROM users WHERE id = $1`

	err := repository.db.QueryRow(query, userID).Scan(
		&user.ID, &user.Username, &user.PasswordHash,
		&user.IsPremium, &user.VerifiedBadge,
		&user.SwipeQuota, &user.QuotaResetTime,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}

	return user, err
}
