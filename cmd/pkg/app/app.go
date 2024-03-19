package app

import (
	"context"
	"fmt"
	"log"
	"net"

	api "github.com/Mehak2716/sample-manager/internal/apis"
	configstore "github.com/Mehak2716/sample-manager/internal/config"
	"github.com/swiggy-private/gocommons/grpc"

	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
)

func Start(ctx context.Context, errch chan<- error) {

	configstore.Initialize()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", configstore.Get().GrpcPort))
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	log.Printf("Listening on %v", configstore.Get().GrpcPort)

	// Setup grpc server
	grpcServer := grpc.NewServer()
	sampleManagerServer := api.SampleManagerServer{}
	samplev1.RegisterSampleManagerServer(grpcServer, sampleManagerServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
