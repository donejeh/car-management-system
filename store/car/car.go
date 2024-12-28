package car

import (
	"context"
	"database/sql"

	"github.com/donejeh/car-management-system/models"
)

type Store struct {
	db *sql.DB
}

// NewStore creates a new store
func NewStore(db *sql.DB) Store {
	return Store{db: db}
}

func (s Store) GetCarById(ctx context.Context, id int) (models.Car, error) {
	// ...
}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) {
	// ...
}

func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	// ...
}

func (s Store) UpdateCar(ctx context.Context, id int, carReq *models.CarRequest) (models.Car, error) {
	// ...
}

func (s Store) DeleteCar(ctx context.Context, id int) (models.Car, error) {
	// ...
}
