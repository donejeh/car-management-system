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
	var car models.Car

	// Query the database
	query := `SELECT c.id, c.name, c.year, brand , c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cyclinders, e.car_range FROM car c LEFT JOIN engine e ON c.engine_id = e.id WHERE c.id=$1`

	row := s.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.ID, &car.Price, &car.CreatedAt, &car.UpdatedAt, &car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange)

	if err != nil {
		if err == sql.ErrNoRows {
			return car, nil
		}

		return car, err
	}

	return car, nil

}

func (s Store) GetCarByBrand(ctx context.Context, brand string, isEngine bool) ([]models.Car, error) {
	var cars []models.Car
	var query string

	if isEngine {
		query = `SELECT c.id, c.name, c.year, brand , c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at, e.id, e.displacement, e.no_of_cyclinders, e.car_range FROM car c LEFT JOIN engine e ON c.engine_id = e.id WHERE brand=$1`
	} else {
		query = `SELECT c.id, c.name, c.year, brand , c.fuel_type, c.engine_id, c.price, c.created_at, c.updated_at FROM car c WHERE brand=$1`
	}

	rows, err := s.db.QueryContext(ctx, query, brand)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var car models.Car

		if isEngine {
			var engine models.Engine
			err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.ID, &car.Price, &car.CreatedAt, &car.UpdatedAt, &car.Engine.EngineID, &car.Engine.Displacement, &car.Engine.NoOfCylinders, &car.Engine.CarRange)
			if err != nil {
				return nil, err
			}

			car.Engine = engine
		} else {
			err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.FuelType, &car.ID, &car.Price, &car.CreatedAt, &car.UpdatedAt)
			if err != nil {
				return nil, err
			}
		}

		cars = append(cars, car)
	}
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
