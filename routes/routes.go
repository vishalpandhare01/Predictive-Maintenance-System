package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/services"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/equipment", services.CreateEquipment)

	app.Post("/sensors", services.AddSensorData)
	app.Post("/maintenance", services.AddMaintenanceLog)

	app.Get("/equipment", services.GetEquipmentList)
	app.Get("/equipment/:id", services.GetEquipmentData)
	app.Delete("/equipment/:id", services.DeleteEquipmentData)

	app.Get("/predictions/:id", services.GetPredictions)
}
