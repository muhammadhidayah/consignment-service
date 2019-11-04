package consignment

import (
	pb "github.com/muhammadhidayah/consignment-service/proto/consignment"
)

type Repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() ([]*pb.Consignment, error)
}
