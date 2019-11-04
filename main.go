package main

import (
	"log"
	"net"

	deliveryGRPC "github.com/muhammadhidayah/consignment-service/consignment/delivery/grpc"
	"github.com/muhammadhidayah/consignment-service/consignment/repository"
	"github.com/muhammadhidayah/consignment-service/consignment/usecase"
	pb "github.com/muhammadhidayah/consignment-service/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":5001"
)

func main() {
	repo := repository.NewSyncRepository()
	consUC := usecase.NewConsignmentUC(repo)
	handler := deliveryGRPC.NewConsignmentHandler(consUC)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterShippingServiceServer(s, handler)

	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to server: %v", err)
	}

}
