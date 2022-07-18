package helpers

import (
	"gitlab.com/propetrov/project_example/pkg/api/public"
	"gitlab.com/propetrov/project_example/pkg/plants/dao"
	"gitlab.com/propetrov/project_example/pkg/plants/dto"
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

func ConvertDTOToPulbicPlant(src *dto.Plant) *public.Plant {
	return &public.Plant{
		Id:   uint32(src.ID),
		Name: src.Name,
	}
}

func ConvertSlicesDTOToPublicPlant(src []*dto.Plant) []*public.Plant {
	dst := make([]*public.Plant, len(src))
	for i := range src {
		dst[i] = &public.Plant{
			Id:   uint32(src[i].ID),
			Name: src[i].Name,
		}
	}
	return dst
}
