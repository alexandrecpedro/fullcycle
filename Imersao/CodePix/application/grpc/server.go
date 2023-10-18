package grpc

import (
	"fmt"
	"github.com/alexandrecpedro/fullcycle/Imersao/CodePix/application/grpc/pb"
	"github.com/alexandrecpedro/fullcycle/Imersao/CodePix/application/usecase"
	"github.com/alexandrecpedro/fullcycle/Imersao/CodePix/infrastructure/repository"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGrpcServer(database *gorm.DB, port int) {
	// COMMENT: Initializes gRPC server
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	pixUseCase := usecase.PixUseCase{PixKeyRepository: pixRepository}
	pixGrpcService := NewPixGrpcService(pixUseCase)
	pb.RegisterPixServiceServer(grpcServer, pixGrpcService)

	address := fmt.Sprintf("0.0.0.0:%d", port)
	listener, err := net.Listen("tcp", address)
	
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}

	log.Printf("gRPC server has been started on port %d", port)
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}