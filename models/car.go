package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Year      string    `json:"year"`
	Brand     string    `json:"brand"`
	FuelType  string    `json:"fuel_type"`
	Engine    Engine    `json:"engine"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ValidateRequest( carReq CarRequest) error {
	if err := validateName(carReq.Name); err != nil {
		return err
	}

	if err := validateYear(carReq.Year); err != nil {
		return err
	}

	if err := validateBrand(carReq.Brand); err != nil {
		return err
	}

	if err := validateFuelType(carReq.FuelType); err != nil {
		return err
	}

	if err := validateEngine(carReq.Engine); err != nil {
		return err
	}

	if err := validatePrice(carReq.Price); err != nil {
		return err
	}

	return nil
	
}

type CarRequest struct {
	Name     string  `json:"name"`
	Year     string  `json:"year"`
	Brand    string  `json:"brand"`
	FuelType string  `json:"fuel_type"`
	Engine   Engine  `json:"engine"`
	Price    float64 `json:"price"`
}

// validateCarRequest validates the CarRequest struct

func validateName(name string) error {
	if name == "" {
		return errors.New("name is required")
	}
	return nil
}

func validateYear(year string) error {
	if year == "" {
		return errors.New("year is required")
	}

	_, error := strconv.Atoi(year)
	if error != nil {
		return errors.New("year must be a valid number")
	}

	currentYear := time.Now().Year()
	yearInt, _ := strconv.Atoi(year)
	if yearInt < 1886 || yearInt > currentYear {
		return errors.New("year must be between 1886 and current year")
	}

	return nil
}

func validateBrand(brand string) error {
	if brand == "" {
		return errors.New("brand is required")
	}
	return nil
}

func validateFuelType(FuelType string) error {

	validateFuelType := []string{"Petrol", "Diesel", "Electric", "Hybrid"}

	for _, fuelType := range validateFuelType {
		if fuelType == FuelType {
			return nil
		}
	}
	return errors.New("fuel type must be Petrol, Diesel, Electric or Hybrid")
}

func validateEngine(engine Engine) error {

	if engine.EngineID == uuid.Nil {
		return errors.New("EngineID is required")
	}

	if engine.Displacement <= 0 {
		return errors.New("displacement is required")
	}

	if engine.NoOfCylinders <= 0 {
		return errors.New("number of cylinders is required")
	}

	if engine.CarRange == 0 {
		return errors.New("car range is required")
	}

	return nil
}

func validatePrice(price float64) error {
	if price <= 0 {
		return errors.New("price is required or greater than 0")
	}
	return nil
}
