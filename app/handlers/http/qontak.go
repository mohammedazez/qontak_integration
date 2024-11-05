package http

import (
	"qontak_integration/app/models"
	"qontak_integration/app/usecase"
	"qontak_integration/pkg/base"
	"qontak_integration/pkg/repository"
	"qontak_integration/pkg/utils"
	"strconv"

	Error "github.com/ewinjuman/go-lib/error"
	"github.com/gofiber/fiber/v2"
)

func GetWebhook(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return ctx.Response(nil, Error.New(repository.BadRequestCode, repository.FailedStatus, "ID is required!"))
	}

	webhook := usecase.NeWebhookUsecase(ctx.Session)
	result, err := webhook.GetWebhook(id)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func ListWebhook(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	webhook := usecase.NeWebhookUsecase(ctx.Session)
	result, err := webhook.ListWebhook()
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func SaveWebhook(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new  struct.
	request := &models.WebhookResult{}

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
	webhook := usecase.NeWebhookUsecase(ctx.Session)
	result, err := webhook.SaveWebhook(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func QontakArchived(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new user auth struct.
	request := &models.ArchivedRequest{}

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
	inbound := usecase.NewInboundUsecase(ctx.Session)
	result, err := inbound.QontakArchived(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func QontakInbound(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new user auth struct.
	request := &models.QontakGeneralMessage{}

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
	inbound := usecase.NewInboundUsecase(ctx.Session)
	result, err := inbound.QontakInbound(request, ctx.Body())
	// Return status 200 OK.
	return ctx.Response(result, err)
}
