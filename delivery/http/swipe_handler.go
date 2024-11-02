package http

import (
	"api-dating/models"
	"api-dating/usecase"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type SwipeHandler struct {
	SwipeUsecase *usecase.SwipeUsecase
}

func NewSwipeHandler(swipeUsecase *usecase.SwipeUsecase) *SwipeHandler {
	return &SwipeHandler{SwipeUsecase: swipeUsecase}
}

func (h *SwipeHandler) Swipe(c *fiber.Ctx) error {
	userID, _ := c.Locals("userID").(uint)
	profileID, _ := strconv.Atoi(c.FormValue("profile_id"))
	direction := c.FormValue("direction")

	err := h.SwipeUsecase.Swipe(userID, profileID, direction)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: err.Error(),
			Data:    []string{},
		})
	}

	return c.JSON(models.Response{
		Message: "Swipe successful",
		Data:    []string{},
	})
}
