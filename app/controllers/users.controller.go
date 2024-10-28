package controllers

import (
	"errors"

	"github.com/gofiber/fiber/v2"

	"github.com/muhammadfikri4/go-plate/app/dto"
	"github.com/muhammadfikri4/go-plate/app/services"
	"github.com/muhammadfikri4/go-plate/utils"
)

type UserController struct {
	service services.UserService
}

func NewUserController(service services.UserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (ctx *UserController) GetUsers(c *fiber.Ctx) error {
	q := new(utils.QueryParams)
	if err := c.QueryParser(q); err != nil {
		return err
	}

	h := &utils.ResponseHandler{}

	users, meta, err := ctx.service.GetAllUsers(*q)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}
	return h.Ok(c, users, "users fetched successfully", &meta)
}

func (ctx *UserController) CreateUser(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	var dto dto.CreateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	err := ctx.service.CreateUser(&dto)
	if err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}
	return h.Created(c, nil, "user created successfully")
}

func (ctx *UserController) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}

	user, err := ctx.service.GetUserById(id)
	if err != nil {
		return h.NotFound(c, []string{err.Error()})
	}

	return h.Ok(c, user, "users fetched successfully", nil)
}

func (ctx *UserController) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}

	var dto *dto.UpdateUserDTO
	if err := c.BodyParser(&dto); err != nil {
		return h.BadRequest(c, []string{err.Error()})
	}

	user, err := ctx.service.UpdateUser(id, dto)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error()})
	}
	return h.Ok(c, user, "user updated successfully", nil)
}

func (ctx *UserController) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	h := &utils.ResponseHandler{}
	err := ctx.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, errors.New("user not found")) {
			return h.NotFound(c, []string{err.Error()})
		}
		return h.InternalServerError(c, []string{err.Error()})
	}

	return h.Ok(c, nil, "User deleted successfully", nil)
}
