package private

import (
	"context"

	"github.com/petrolax/project-template/pkg/api/private"
	"github.com/petrolax/project-template/pkg/transport/helpers"
)

func (c *Controller) GetPlants(ctx context.Context, req *private.GetPlantsRequest) (*private.GetPlantsResponse, error) {
	c.lg.Info("GetPlantsHandler")

	plants, err := c.service.GetPlants(ctx)
	if err != nil {
		return nil, err
	}
	return &private.GetPlantsResponse{Plants: helpers.ConvertSlicesDTOToPrivatePlant(plants)}, nil
}
