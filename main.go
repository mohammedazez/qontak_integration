package main

import (
	"github.com/common-nighthawk/go-figure"
	Logger "github.com/ewinjuman/go-lib/logger"
	Session "github.com/ewinjuman/go-lib/session"
	"github.com/gofiber/fiber/v2"
	"qontak_integration/app/consumer"
	"qontak_integration/pkg/configs"
	"qontak_integration/pkg/middleware"
	"qontak_integration/pkg/routes"
	"qontak_integration/pkg/utils"
)

// @title Skeleton Service API
// @version 2.0
// @description Skeleton service using golang and fiber framework.
// @Still continuing to develop
func main() {
	// Define Fiber config.
	config := configs.FiberConfig()

	// Define a new Fiber app with config.
	app := fiber.New(config)

	// Middlewares.
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes.
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.PrivateRoutes(app) // Register a private_libs.sh routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	myFigure := figure.NewColorFigure("eways - qontak_integration", "", "green", true)
	myFigure.Print()

	go consumer.OutboundBroadcastConsumer(Session.New(Logger.New(configs.Config.Logger)).
		SetInstitutionID(configs.Config.Apps.Name).
		SetAppName(configs.Config.Apps.Name).
		SetURL("/start-broadcast-consumer").
		SetMethod("QUEUE"))

	go consumer.OutboundSendMessageConsumer(Session.New(Logger.New(configs.Config.Logger)).
		SetInstitutionID(configs.Config.Apps.Name).
		SetAppName(configs.Config.Apps.Name).
		SetURL("/start-send-message-consumer").
		SetMethod("QUEUE"))

	// Start server (with or without graceful shutdown).
	if configs.Config.Apps.Mode == "local" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
