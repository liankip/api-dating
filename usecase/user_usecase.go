package usecase

import (
	"api-dating/models"
	"api-dating/repository"
	"api-dating/utils"
	"errors"
	"time"
)

type UserUsecase struct {
	UserRepository           repository.UserRepository
	PremiumPackageRepository repository.PremiumPackageRepository
}

func NewUserUsecase(userRepository repository.UserRepository, premiumPackageRepository repository.PremiumPackageRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository:           userRepository,
		PremiumPackageRepository: premiumPackageRepository,
	}
}

func (userUseCase *UserUsecase) Signup(username, password string) error {
	_, err := userUseCase.UserRepository.FindUserByUsername(username)

	if err == nil {
		return errors.New("user already exists")
	}

	passwordHash, _ := utils.HashPassword(password)

	user := &models.User{
		Username:       username,
		PasswordHash:   passwordHash,
		SwipeQuota:     10,
		QuotaResetTime: time.Now().Add(24 * time.Hour),
	}

	_ = models.Profile{
		UserID:   user.ID,
		Username: user.Username,
	}

	return userUseCase.UserRepository.CreateUser(user)
}

func (userUseCase *UserUsecase) Signin(username, password string) (*models.User, error) {
	user, err := userUseCase.UserRepository.FindUserByUsername(username)

	if err != nil || !utils.CheckPasswordHash(password, user.PasswordHash) { // implement checkPasswordHash
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}

func (userUseCase *UserUsecase) PremiumActive(userID uint, packageID int) error {
	premiumPackage, err := userUseCase.PremiumPackageRepository.GetPremiumPackageByID(packageID)
	if err != nil {
		return errors.New("premium package not found")
	}

	err = userUseCase.PremiumPackageRepository.ActivateUserPremium(userID, premiumPackage.ID)
	if err != nil {
		return errors.New("failed to activate premium package: " + err.Error())
	}

	err = userUseCase.UserRepository.UpdateSwipeQuota(userID, int(premiumPackage.Quota))
	if err != nil {
		return errors.New("failed to update user swipe quota: " + err.Error())
	}

	return nil
}
