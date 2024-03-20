package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Mehak2716/sample-manager/internal/config"
	"github.com/Mehak2716/sample-manager/internal/interceptors"
	"github.com/Mehak2716/sample-manager/internal/repository"
	server "github.com/Mehak2716/sample-manager/internal/server"
	"github.com/Mehak2716/sample-manager/internal/services"
	"google.golang.org/grpc"

	samplev1 "github.com/Mehak2716/sample-manager-proto/v1"
)

const grpcPort uint32 = 8090

func Start(ctx context.Context, errch chan<- error) {

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", grpcPort))
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}
	log.Printf("Listening on %v", grpcPort)

	// Setup grpc server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(interceptors.ValidateRequestHandler),
	)
	db := config.DatabaseConnection()
	sampleRepo := repository.SampleRepository{DB: db}
	sampleService := services.SampleService{Repo: sampleRepo}
	sampleManagerServer := server.NewSampleManagerServer(sampleService)

	samplev1.RegisterSampleManagerServer(grpcServer, sampleManagerServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
