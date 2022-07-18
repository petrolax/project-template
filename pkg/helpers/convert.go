package helpers

import (
	"github.com/petrolax/project-template/pkg/plants/dao"
	"github.com/petrolax/project-template/pkg/plants/dto"
)

func ConvertSlicesDAOToDTO(src []*dao.Plant) []*dto.Plant {
	dst := make([]*dto.Plant, len(src))
	for i := range src {
		dst[i] = &dto.Plant{
			ID:   src[i].ID,
			Name: src[i].Name,
		}
	}
	return dst
}
