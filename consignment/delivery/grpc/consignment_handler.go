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

func (handler *ConsignmentHandler) CreateConsignment(ctx context.Context, req *pb.Consignment) (*pb.Response, error) {
	consignment, err := handler.CUsecase.Create(req)

	if err != nil {
		return nil, err
	}

	return &pb.Response{Created: true, Consignment: consignment}, nil
}
