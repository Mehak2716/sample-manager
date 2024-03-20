package server

import (
	"context"

	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
	services "github.com/Mehak2716/sample-manager/internal/services"
)

type SampleManagerServer struct {
	sampleService services.SampleService
	samplev1.SampleManagerServer
}

func NewSampleManagerServer(sampleService services.SampleService) *SampleManagerServer {
	return &SampleManagerServer{
		sampleService: sampleService,
	}
}

func (server *SampleManagerServer) CreateSampleMapping(ctx context.Context, req *samplev1.SampleMappingRequest) (*samplev1.SampleMappingResponse, error) {

	return server.sampleService.CreateMapping(req)
}

func (server *SampleManagerServer) GetSampleSKUs(ctx context.Context, req *samplev1.GetSampleIDsRequest) (*samplev1.GetSampleIDsResponse, error) {

	return server.sampleService.GetSampleSKUs(req)
}
