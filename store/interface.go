package store

import (
	"context"

	"github.com/donejeh/car-management-system/models"
)

type CarStoreInterface interface {
	GetCarById(ctx context.Context, id string) (*models.Car, error)
	GetAllCars(ctx context.Context) ([]models.Car, error)
	CreateCar(ctx context.Context, car *models.Car) error
	UpdateCar(ctx context.Context, car *models.Car) error
	DeleteCar(ctx context.Context, id string) error
	GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error)
}
