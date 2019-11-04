package grpc

import (
	"context"

	"github.com/muhammadhidayah/consignment-service/consignment"
	pb "github.com/muhammadhidayah/consignment-service/proto/consignment"
)

type ConsignmentHandler struct {
	CUsecase consignment.Usecase
}

func NewConsignmentHandler(consusecase consignment.Usecase) *ConsignmentHandler {
	return &ConsignmentHandler{consusecase}
}

func (handler *ConsignmentHandler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {
	consignment, err := handler.CUsecase.Create(req)

	if err != nil {
		return err
	}

	res.Created = true
	res.Consignment = consignment
	return nil
}

func (handler *ConsignmentHandler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments, _ := handler.CUsecase.GetAll()
	res.Consignments = consignments
	return nil
}
