package mapper

import (
	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
	"github.com/Mehak2716/sample-manager/internal/models"
)

func MapToSampleMapping(req *samplev1.SampleMappingRequest) models.SampleMapping {

	return models.SampleMapping{
		CustomerSegment: req.CustomerSegment,
		ProductID:       req.ProductID,
		SampleProductID: req.SampleProductID,
	}
}

func MapToSampleMappingResponse(sampleMapping *models.SampleMapping) *samplev1.SampleMappingResponse {

	return &samplev1.SampleMappingResponse{
		ID:              int64(sampleMapping.ID),
		CustomerSegment: sampleMapping.CustomerSegment,
		ProductID:       sampleMapping.ProductID,
		SampleProductID: sampleMapping.SampleProductID,
	}
}
