package usecase

import (
	"api-dating/models"
	"api-dating/repository"
	"errors"
	"time"
)

type SwipeUsecase struct {
	UserRepository  repository.UserRepository
	SwipeRepository repository.SwipeRepository
}

func NewSwipeUsecase(userRepository repository.UserRepository, swipeRepository repository.SwipeRepository) *SwipeUsecase {
	return &SwipeUsecase{
		UserRepository:  userRepository,
		SwipeRepository: swipeRepository,
	}
}

func (uc *SwipeUsecase) Swipe(userID uint, profileID int, direction string) error {
	hasSwiped, err := uc.SwipeRepository.HasSwipedToday(userID, profileID)
	if err != nil {
		return err
	}

	if hasSwiped {
		return errors.New("cannot swipe the same profile twice in a day")
	}

	user, err := uc.UserRepository.GetUserByID(userID)
	if err != nil {
		return err
	}

	if user.ID == profileID {
		return errors.New("cannot swipe yourself")
	}

	if time.Now().After(user.QuotaResetTime) {
		user.SwipeQuota = 10
		user.QuotaResetTime = time.Now().Add(24 * time.Hour)
		err = uc.UserRepository.UpdateSwipeQuota(userID, user.SwipeQuota)
		if err != nil {
			return err
		}
	}

	if user.SwipeQuota <= 0 {
		return errors.New("daily swipe quota exceeded")
	}

	user.SwipeQuota -= 1
	err = uc.UserRepository.UpdateSwipeQuota(userID, user.SwipeQuota)
	if err != nil {
		return err
	}

	swipe := &models.Swipe{
		UserID:    uint(userID),
		ProfileID: uint(profileID),
		Direction: direction,
		CreatedAt: time.Now(),
	}

	return uc.SwipeRepository.RecordSwipe(swipe)
}
