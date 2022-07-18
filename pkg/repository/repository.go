package repository

import (
	"context"

	"gitlab.com/propetrov/project_example/pkg/plants/dao"
)

type Repository interface {
	AddPlant(ctx context.Context, newPlant *dao.Plant) error
	GetPlantByName(ctx context.Context, name string) (*dao.Plant, error)
	GetPlants(ctx context.Context) ([]*dao.Plant, error)
	UpdatePlant(ctx context.Context, plant *dao.Plant) error
	DeletePlant(ctx context.Context, plant dao.Plant) error
}
