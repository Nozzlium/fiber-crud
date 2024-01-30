package controller

import (
	"fiber-crud/entity"
	"fiber-crud/param"
	"fiber-crud/requestbody"
	"fiber-crud/service"

	"github.com/gofiber/fiber/v2"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(
	userService service.UserService,
) *UserControllerImpl {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) Create(ctx *fiber.Ctx) error {
	requestBody := requestbody.CreateUser{}
	err := ctx.BodyParser(&requestBody)
	if err != nil {
		return err
	}
	result, err := controller.UserService.Create(
		ctx.Context(),
		entity.User{
			Name:        requestBody.Name,
			PhoneNumber: requestBody.PhoneNumber,
		},
	)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   result,
	})
}

func (controller *UserControllerImpl) Get(ctx *fiber.Ctx) error {
	page := ctx.QueryInt("page", 1)
	pageNo := ctx.QueryInt("pageSize", 5)
	results, err := controller.UserService.Get(
		ctx.Context(),
		param.UserParam{
			PageNo:   page,
			PageSize: pageNo,
		},
	)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data": fiber.Map{
			"page":     page,
			"pageSize": len(results),
			"users":    results,
		},
	})
}

func (controller *UserControllerImpl) GetById(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}
	result, err := controller.UserService.GetById(
		ctx.Context(),
		entity.User{
			ID: int32(id),
		},
	)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   result,
	})
}

func (controller *UserControllerImpl) Update(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}
	requestBody := requestbody.UpdateUser{}
	ctx.BodyParser(&requestBody)
	result, err := controller.UserService.Update(
		ctx.Context(),
		entity.User{
			ID:          int32(id),
			Name:        requestBody.Name,
			PhoneNumber: requestBody.PhoneNumber,
		},
	)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   result,
	})
}

func (controller *UserControllerImpl) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id", 0)
	if err != nil {
		return err
	}
	result, err := controller.UserService.Delete(
		ctx.Context(),
		entity.User{
			ID: int32(id),
		},
	)
	if err != nil {
		return err
	}
	return ctx.JSON(fiber.Map{
		"code":   fiber.StatusOK,
		"status": "OK",
		"data":   result,
	})
}
