package http

import (
	Error "github.com/ewinjuman/go-lib/error"
	"github.com/gofiber/fiber/v2"
	"qontak_integration/app/models"
	"qontak_integration/app/usecase"
	"qontak_integration/pkg/base"
	"qontak_integration/pkg/repository"
	"qontak_integration/pkg/utils"
	"strconv"
)

func GetWaTemplate(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	clientID, err := strconv.Atoi(c.Query("client_id"))
	if err != nil || clientID <= 0 {
		return ctx.Response(nil, Error.NewError(repository.BadRequestCode, repository.FailedStatus, "Client ID is required!"))
	}
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 500
		err = nil
	}

	offset, err := strconv.Atoi(c.Query("offset"))
	if err != nil {
		offset = 0
		err = nil
	}

	template := usecase.NewTemplateUsecase(ctx.Session)
	res, err := template.GetTemplate(&models.GetTemplateRequest{
		ClientID: clientID,
		Offset:   offset,
		Limit:    limit,
	})

	return ctx.Response(res, err)
}

func SendWaMessage(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new user auth struct.
	request := &models.SendWhatsappRequest{}

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
	result, err := outbound.SendWaMessage(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func SendWaBroadcastdirect(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new user auth struct.
	request := &models.BroadcastDirectRequest{}

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
	result, err := outbound.SendWaBroadCastDirect(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}

func SendWaBroadcastdirectBulk(c *fiber.Ctx) error {
	ctx := base.NewContext(c)

	// Create a new user auth struct.
	request := &models.BroadcastDirectBulkRequest{}

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
	result, err := outbound.SendWaBroadCastDirectBulk(request)
	// Return status 200 OK.
	return ctx.Response(result, err)
}
