package interceptors

import (
	"context"
	"log"

	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateRequestHandler(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("Intercepting unary request: %v\n", info.FullMethod)

	switch info.FullMethod {
	case "/SampleManager/CreateSampleMapping":

		r := req.(*samplev1.SampleMappingRequest)
		if r.CustomerSegment == "" || r.ProductID == "" || r.SampleProductID == "" {
			return nil, status.Error(codes.InvalidArgument, "Customer Segment, ProductID, SampleProductID must be provided")
		}

	case "/SampleManager/GetSampleIDs":

		r := req.(*samplev1.GetSampleIDsRequest)
		if r.CustomerSegments == nil || r.ProductIDs == nil {
			return nil, status.Error(codes.InvalidArgument, "Provide Customer Segments and ProductIDs")
		}

	}
	resp, err := handler(ctx, req)
	return resp, err

}
