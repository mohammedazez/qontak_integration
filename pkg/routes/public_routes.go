package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	httpHandler "qontak_integration/app/handlers/http"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api")
	v1 := route.Group("/v1")
	v1.Get("/monitor", monitor.New(monitor.Config{Title: "Qontak Integration Metrics Page"}))
	v1.Get("/health", httpHandler.HealthCheck) // register a new user

	v1.Get("/templates/whatsapp", httpHandler.GetWaTemplate)

	whatsapp := v1.Group("/whatsapp")

	whatsapp.Post("/message", httpHandler.SendWaMessage)
	whatsapp.Post("/broadcast/direct", httpHandler.SendWaBroadcastdirect)
	whatsapp.Post("/broadcast/direct/bulk", httpHandler.SendWaBroadcastdirectBulk)

	instagram := v1.Group("/instagram")
	instagram.Post("/message", httpHandler.SendInstagramMessage)

	qontak := v1.Group("qontak")
	qontak.Post("/inbound", httpHandler.QontakInbound)
	qontak.Post("/archived", httpHandler.QontakArchived)

	v1.Get("/webhook/:id", httpHandler.GetWebhook)
	v1.Get("/webhooks", httpHandler.ListWebhook)
	v1.Post("/webhook", httpHandler.SaveWebhook)
}
