package usecase

import (
	"github.com/muhammadhidayah/consignment-service/consignment"
	pb "github.com/muhammadhidayah/consignment-service/proto/consignment"
)

type consignmentUC struct {
	repo consignment.Repository
}

func NewConsignmentUC(repo consignment.Repository) consignment.Usecase {
	return &consignmentUC{repo}
}

func (usecase *consignmentUC) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	res, err := usecase.repo.Create(consignment)

	if err != nil {
		return nil, err
	}

	return res, nil
}
