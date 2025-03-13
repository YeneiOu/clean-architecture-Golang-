package controllers

import (
	"bank/helper"
	"bank/modules/entities"
	"github.com/gofiber/fiber/v2"
)

type userController struct {
	UserUse entities.UserUsecase
}

func NewUsersController(r fiber.Router, userUse entities.UserUsecase) {
	controller := &userController{
		UserUse: userUse,
	}
	r.Get("/", controller.GetUsers)
}

func (controller *userController) GetUsers(c *fiber.Ctx) error {

	res, err := controller.UserUse.GetUsers()

	if err != nil {
		return helper.RespondWithJSON(
			c,
			fiber.ErrBadRequest.Message,
			fiber.StatusBadRequest,
			"Some thing error",
			[]string{},
		)
	}

	if len(res) == 0 {
		return helper.RespondWithJSON(
			c,
			"OK",
			fiber.StatusOK,
			"No users",
			[]string{},
		)
	}
	return helper.RespondWithJSON(
		c,
		"OK",
		fiber.StatusOK,
		"Success",
		res,
	)
}
