package main

import (
	"fmt"

	"github.com/micro/go-micro"
	deliveryGRPC "github.com/muhammadhidayah/consignment-service/consignment/delivery/grpc"
	"github.com/muhammadhidayah/consignment-service/consignment/repository"
	"github.com/muhammadhidayah/consignment-service/consignment/usecase"
	pb "github.com/muhammadhidayah/consignment-service/proto/consignment"
)

func main() {
	repo := repository.NewSyncRepository()
	consUC := usecase.NewConsignmentUC(repo)
	handler := deliveryGRPC.NewConsignmentHandler(consUC)

	srv := micro.NewService(
		micro.Name("consignment"),
	)

	srv.Init()

	pb.RegisterShippingServiceHandler(srv.Server(), handler)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
