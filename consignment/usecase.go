package consignment

import pb "github.com/muhammadhidayah/consignment-service/proto/consignment"

type Usecase interface {
	Create(consignment *pb.Consignment) (*pb.Consignment, error)
}
