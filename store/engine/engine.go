package engine

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/donejeh/car-management-system/models"
	"github.com/google/uuid"
)

type EngineStore struct {
	db *sql.DB
}

func NewEngineStore(db *sql.DB) *EngineStore {
	return &EngineStore{db: db}
}

func (e *EngineStore) EngineById(ctx context.Context, id string) (models.Engine, error) {

	var engine models.Engine

	tx, err := e.db.BeginTx(ctx, nil)

	if err != nil {
		return engine, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v\n", rbErr)
			}
		} else {
			if CmErr := tx.Commit(); CmErr != nil {
				fmt.Printf("CmError committing transaction: %v\n", CmErr)
			}
		}
	}()

	err = tx.QueryRowContext(ctx, "SELECT id, displacement, no_of_cylinders, car_range FROM engine WHERE id=$1", id).Scan(&engine.EngineID, &engine.Displacement, &engine.NoOfCylinders, &engine.CarRange)

	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return engine, nil
		}

		return engine, err
	}

	return engine, err
}

func (e *EngineStore) EngineCreate(ctx context.Context, engineReq *models.EngineRequest) (models.Engine, error) {

	tx, err := e.db.BeginTx(ctx, nil)

	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v\n", rbErr)
			}
		} else {
			if CmErr := tx.Commit(); CmErr != nil {
				fmt.Printf("Error committing transaction: %v\n", CmErr)
			}
		}
	}()

	engineId := uuid.New()

	_, err = tx.ExecContext(ctx, "INSERT INTO engine (id, displacement, no_of_cylinders, car_range) VALUES ($1, $2, $3, $4)", engineId, engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange)

	if err != nil {
		return models.Engine{}, err
	}

	engine := models.Engine{
		EngineID:      engineId,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e *EngineStore) EngineUpdate(ctx context.Context, id string, engineReq *models.EngineRequest) (models.Engine, error) {

	engineID, err := uuid.Parse(id)

	if err != nil {
		return models.Engine{}, fmt.Errorf("invalid engine id: %w", err)
	}

	tx, err := e.db.BeginTx(ctx, nil)

	if err != nil {
		return models.Engine{}, err
	}

	defer func() {
		if err != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				fmt.Printf("Error rolling back transaction: %v\n", rbErr)
			}
		} else {
			if CmErr := tx.Commit(); CmErr != nil {
				fmt.Printf("Error committing transaction: %v\n", CmErr)
			}
		}
	}()

	result, error := tx.ExecContext(ctx, "UPDATE engine SET displacement=$1, no_of_cylinders=$2, car_range=$3 WHERE id=$4", engineReq.Displacement, engineReq.NoOfCylinders, engineReq.CarRange, engineID)

	if error != nil {
		return models.Engine{}, error
	}

	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 {
		return models.Engine{}, errors.New("no rows affected")
	}

	engine := models.Engine{
		EngineID:      engineID,
		Displacement:  engineReq.Displacement,
		NoOfCylinders: engineReq.NoOfCylinders,
		CarRange:      engineReq.CarRange,
	}

	return engine, nil
}

func (e *EngineStore) DeleteEngine(ctx context.Context, id string) (models.Engine, error) {
	// ...
}
