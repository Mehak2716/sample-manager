package services

import (
	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
	"github.com/Mehak2716/sample-manager/internal/mapper"
	"github.com/Mehak2716/sample-manager/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SampleService struct {
	Repo repository.SampleRepository
}

func (service *SampleService) CreateMapping(req *samplev1.SampleMappingRequest) (*samplev1.SampleMappingResponse, error) {

	sampleMapping := mapper.MapToSampleMapping(req)

	if service.Repo.IsExists(sampleMapping.CustomerSegment, sampleMapping.ProductID) {
		return nil, status.Errorf(codes.AlreadyExists, "Product for this customer segment is already mapped to a sample")
	}

	createdSampleMapping, err := service.Repo.Save(&sampleMapping)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create sample mapping")
	}

	response := mapper.MapToSampleMappingResponse(createdSampleMapping)
	return response, nil
}

func (service *SampleService) GetSampleIDs(req *samplev1.GetSampleIDsRequest) (*samplev1.GetSampleIDsResponse, error) {

	segments := req.CustomerSegments
	productIDS := req.ProductIDs

	sampleMappings, err := service.Repo.FetchSampleIDs(segments, productIDS)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to get samples")
	}

	response := mapper.MapToSampleIDSResponse(sampleMappings)
	return response, nil

}
