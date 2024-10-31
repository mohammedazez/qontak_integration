package http

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/gofiber/fiber/v2"
	"qontak_integration/app/models"
	"qontak_integration/app/usecase"
	"qontak_integration/pkg/base"
	"qontak_integration/pkg/repository"
	"qontak_integration/pkg/utils"
)

func SendInstagramMessage(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new user auth struct.
	request := &models.SendInstagramRequest{}

	// Checking received data from JSON body.
	if err := ctx.BodyParser(request); err != nil {
		// Return status 400 and error message.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}

	// Create a new validator for a User model.
	validate := utils.NewValidator()
	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		// Return, if some fields are not valid.
		return ctx.Response(nil, Error.New(fiber.StatusBadRequest, repository.FailedStatus, err.Error()))
	}
	outbound := usecase.NewOutboundUsecase(ctx.Session)
	result, err := outbound.SendInstagramMessage(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}
