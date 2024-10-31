package routes

import (
	"github.com/gofiber/fiber/v2"
	httpHandler "qontak_integration/app/handlers/http"
	"qontak_integration/pkg/middleware"
)

// PrivateRoutes func for describe group of private_libs.sh routes.
func PrivateRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	route.Delete("/user/sign/out", middleware.JWTProtected(), httpHandler.GetWaTemplate) // de-authorization user
}
