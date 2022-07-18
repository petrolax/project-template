package private

import (
	"net/http"

	"github.com/gorilla/mux"
	"gitlab.com/propetrov/project_example/pkg/api/private"
	"gitlab.com/propetrov/project_example/pkg/plants"
	"go.uber.org/zap"
)

type Controller struct {
	lg      *zap.Logger
	service plants.Service
	private.UnimplementedPlantsApiServer
}

func NewController(lg *zap.Logger, service plants.Service) *Controller {
	return &Controller{
		lg:      lg,
		service: service,
	}
}

func (c *Controller) Endpoints() http.Handler {
	r := mux.NewRouter()

	return r
}
