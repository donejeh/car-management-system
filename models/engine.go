package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"no_of_cylinders"`
	CarRange      int64     `json:"car_range"`
}

// ValidateEngine validates the engine struct
type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"no_of_cylinders"`
	CarRange      int64 `json:"car_range"`
}

func validateEngineRequest(engine EngineRequest) error {

	err := validateDisplacement(engine.Displacement)
	if err != nil {
		return err
	}

	err = validateNoOfCylinders(engine.NoOfCylinders)
	if err != nil {
		return err
	}

	err = validateCarRange(engine.CarRange)
	if err != nil {
		return err
	}

	return nil
}

// ValidateEngine validates the engine struct
func validateDisplacement(displacement int64) error {

	if displacement <= 0 {
		return errors.New("displacement is required")
	}

	return nil
}

func validateNoOfCylinders(noOfCylinders int64) error {

	if noOfCylinders <= 0 {
		return errors.New("number of cylinders is required or must be greater than 0")
	}

	return nil
}

func validateCarRange(carRange int64) error {

	if carRange <= 0 {
		return errors.New("car range is required or must be greater than 0")
	}

	return nil
}
