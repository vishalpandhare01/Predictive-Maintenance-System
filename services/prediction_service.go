// services/prediction_service.go
package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/db"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/models"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/utils"
)

// Simulate adding sensor data
func AddSensorData(c *fiber.Ctx) error {
	var sensor models.Sensor
	if err := c.BodyParser(&sensor); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	// Save to database
	if err := db.DB.Create(&sensor).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to save data"})
	}
	_, err := utils.PredictMaintenance(sensor.EquipmentID)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(201).JSON(sensor)
}

// create new equipment
func CreateEquipment(c *fiber.Ctx) error {
	var equipment models.Equipment
	if err := c.BodyParser(&equipment); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	if err := db.DB.Create(&equipment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to save data"})
	}
	return c.Status(201).JSON(equipment)

}

// add maintencance log
func AddMaintenanceLog(c *fiber.Ctx) error {
	var maintenanceLog models.MaintenanceLog
	if err := c.BodyParser(&maintenanceLog); err != nil {
		return c.Status(400).JSON(fiber.Map{"message": err.Error()})
	}

	if err := db.DB.Create(&maintenanceLog).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(201).JSON(maintenanceLog)

}

// get equipment list
func GetEquipmentList(c *fiber.Ctx) error {
	var equipment []models.Equipment

	if err := db.DB.Find(&equipment).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}
	return c.Status(200).JSON(equipment)

}

// get equipment
func GetEquipmentData(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment models.Equipment

	if err := db.DB.Where("id = ?", id).First(&equipment).Error; err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"message": "equipment not found"})
		}
		return c.Status(500).JSON(fiber.Map{"message": "Failed Get data"})
	}

	return c.Status(200).JSON(equipment)

}

func DeleteEquipmentData(c *fiber.Ctx) error {
	id := c.Params("id")
	var equipment models.Equipment

	// Find the equipment by ID
	if err := db.DB.Where("id = ?", id).First(&equipment).Error; err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"message": "equipment not found"})
		}
		return c.Status(500).JSON(fiber.Map{"message": "Failed to get data"})
	}

	// Delete dependent records in predictions table
	if err := db.DB.Where("equipment_id = ?", equipment.ID).Delete(&models.Prediction{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete dependent records"})
	}

	// Delete dependent records in sensor table
	if err := db.DB.Where("equipment_id = ?", equipment.ID).Delete(&models.Sensor{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": "Failed to delete dependent records"})
	}

	// Delete the equipment record
	if err := db.DB.Where("id = ?", equipment.ID).Delete(&models.Equipment{}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"message": err.Error()})
	}

	// Return success response
	return c.Status(200).JSON(fiber.Map{
		"message": "Equipment deleted successfully",
	})
}

// get prediction equipment
func GetPredictions(c *fiber.Ctx) error {
	id := c.Params("id")
	var prediction []models.Prediction

	if err := db.DB.Where("equipment_id = ?", id).Preload("Equipment").Find(&prediction).Error; err != nil {
		if err.Error() == "record not found" {
			return c.Status(404).JSON(fiber.Map{"message": "prediction not found"})
		}
		return c.Status(500).JSON(fiber.Map{"message": "Failed Get data"})
	}

	return c.Status(200).JSON(prediction)

}
