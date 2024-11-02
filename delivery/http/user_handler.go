package http

import (
	"api-dating/models"
	"api-dating/usecase"
	"api-dating/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	UserUsecase *usecase.UserUsecase
}

func NewUserHandler(userUsecase *usecase.UserUsecase) *UserHandler {
	return &UserHandler{UserUsecase: userUsecase}
}

func (h *UserHandler) Signup(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	if err := h.UserUsecase.Signup(username, password); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: err.Error(),
			Data:    []string{},
		})
	}

	return c.Status(fiber.StatusCreated).JSON(models.Response{
		Message: "Signup successful",
		Data:    []string{},
	})
}

func (h *UserHandler) Signin(c *fiber.Ctx) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := h.UserUsecase.Signin(username, password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Message: err.Error(),
			Data:    []string{},
		})
	}

	token, err := utils.GenerateJWT(uint(user.ID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: "Failed to generate token",
			Data:    []string{},
		})
	}

	combinedData := fiber.Map{
		"user":  user,
		"token": token,
	}

	return c.JSON(models.Response{
		Message: "Signin successful",
		Data:    combinedData,
	})
}

func (h *UserHandler) PremiumActive(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(uint)
	if !ok || userID == 0 {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: "Invalid or missing user ID",
			Data:    []string{},
		})
	}

	packageID, _ := strconv.Atoi(c.FormValue("package_id"))
	if err := h.UserUsecase.PremiumActive(userID, packageID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
			Message: err.Error(),
			Data:    []string{},
		})
	}

	return c.JSON(models.Response{
		Message: "Premium package activated successfully",
		Data:    map[string]interface{}{"package_id": packageID},
	})
}
