package utils

import (
	"fmt"
	"sort"
	"time"

	"github.com/vishalpandhare01/Predictive-Maintenance-System/db"
	"github.com/vishalpandhare01/Predictive-Maintenance-System/models"
)

func calculateMedian(sensors []models.Sensor) float64 {
	// Extract sensor values
	var values []float64
	for _, sensor := range sensors {
		values = append(values, sensor.Value)
	}

	// Sort the values
	sort.Float64s(values)

	// Calculate median
	n := len(values)
	if n%2 == 0 {
		// If even, return the average of the middle two elements
		return (values[n/2-1] + values[n/2]) / 2.0
	}
	// If odd, return the middle element
	return values[n/2]
}

func PredictMaintenance(equipmentId string) (*models.Prediction, error) {
	// Collect historical sensor data for the specific equipment
	var sensors []models.Sensor
	if err := db.DB.Where("equipment_id = ?", equipmentId).Find(&sensors).Error; err != nil {
		return nil, err
	}

	// If no sensor data exists, we can't predict anything
	if len(sensors) == 0 {
		return nil, fmt.Errorf("no sensor data available for equipment %s", equipmentId)
	}

	// Use median instead of mean to avoid the influence of outliers
	medianValue := calculateMedian(sensors)

	// Debugging: Print out the median value
	fmt.Printf("Calculated Median Value for Equipment %s: %f\n", equipmentId, medianValue)

	// Predict failure probability based on median sensor value
	probability := 0.0
	if medianValue > 70.0 { // Threshold for failure prediction
		probability = 0.8 // High probability of failure
	} else {
		// Debugging: print that the median is below the threshold
		fmt.Printf("Median value %f is below threshold. Probability remains 0.\n", medianValue)
	}

	// Debugging: Print out the failure prediction
	fmt.Printf("Predicted Failure Probability for Equipment %s: %f\n", equipmentId, probability)

	// Create and save prediction record
	prediction := models.Prediction{
		EquipmentID:          equipmentId,
		PredictedFailureDate: time.Now().Add(30 * time.Hour), // Example: Predict failure in 30 hours
		FailureProbability:   probability,
	}

	if err := db.DB.Create(&prediction).Error; err != nil {
		return nil, err
	}

	// Return the prediction
	return &prediction, nil
}
