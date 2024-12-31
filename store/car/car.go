package car

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/donejeh/car-management-system/models"
	"github.com/google/uuid"
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
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func (s Store) CreateCar(ctx context.Context, carReq *models.CarRequest) (models.Car, error) {
	var createdCar models.Car
	var EngineID uuid.UUID

	err := s.db.QueryRowContext(ctx, "SELECT id FROM engine WHERE id=$1", carReq.Engine.EngineID).Scan(&EngineID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return createdCar, errors.New("engine not found")
		}
		return createdCar, err
	}

	carId := uuid.New()
	createdAt := time.Now()
	updatedAt := createdAt

	newCar := models.Car{
		ID:        carId,
		Name:      carReq.Name,
		Year:      carReq.Year,
		Brand:     carReq.Brand,
		FuelType:  carReq.FuelType,
		Engine:    carReq.Engine,
		Price:     carReq.Price,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return createdCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	query := `INSERT INTO car (id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query, newCar.ID, newCar.Name, newCar.Year, newCar.Brand, newCar.FuelType, newCar.Engine.EngineID, newCar.Price, newCar.CreatedAt, newCar.UpdatedAt).Scan(&createdCar.ID, &createdCar.Name, &createdCar.Year, &createdCar.Brand, &createdCar.FuelType, &createdCar.Engine.EngineID, &createdCar.Price, &createdCar.CreatedAt, &createdCar.UpdatedAt)

	if err != nil {
		return createdCar, err
	}

	return createdCar, nil
}

func (s Store) UpdateCar(ctx context.Context, id string, carReq *models.CarRequest) (models.Car, error) {
	var updatedCar models.Car

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return updatedCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	query := `UPDATE car SET name=$2, year=$3, brand=$4, fuel_type=$5, engine_id=$6, price=$7, updated_at=$8 WHERE id=$1 RETURNING id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query, id, carReq.Name, carReq.Year, carReq.Brand, carReq.FuelType, carReq.Engine.EngineID, carReq.Price, time.Now()).Scan(&updatedCar.ID, &updatedCar.Name, &updatedCar.Year, &updatedCar.Brand, &updatedCar.FuelType, &updatedCar.Engine.EngineID, &updatedCar.Price, &updatedCar.CreatedAt, &updatedCar.UpdatedAt)

	if err != nil {
		return updatedCar, err
	}

	return updatedCar, nil
}


func (s Store) DeleteCar(ctx context.Context, id int) (models.Car, error) {
	var deleteCar models.Car

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return deleteCar, err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		err = tx.Commit()
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, name, year, brand, fuel_type, engine_id, price, created_at, updated_at FROM car WHERE id=$1", id).Scan(&deleteCar.ID, &deleteCar.Name, &deleteCar.Year, &deleteCar.Brand, &deleteCar.FuelType, &deleteCar.Engine.EngineID, &deleteCar.Price, &deleteCar.CreatedAt, &deleteCar.UpdatedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Car{}, errors.New("car not found")
		}

		return models.Car{}, err
	}

	result, err := tx.ExecContext(ctx, "DELETE FROM car WHERE id = $1", id)

	if err != nil {
		return models.Car{}, err

	}

	rowAffected, err := result.RowsAffected()

	if err != nil {
		return models.Car{}, err
	}

	if rowAffected == 0 {
		return models.Car{}, errors.New("No row were deleted")
	}

	return deleteCar, nil
}
