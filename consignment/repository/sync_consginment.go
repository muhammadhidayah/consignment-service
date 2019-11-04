package repository

import (
	"sync"

	"github.com/muhammadhidayah/consignment-service/consignment"
	pb "github.com/muhammadhidayah/consignment-service/proto/consignment"
)

type syncRepository struct {
	mu          sync.RWMutex
	consignment []*pb.Consignment
}

func NewSyncRepository() consignment.Repository {
	return &syncRepository{}
}

func (repo *syncRepository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	repo.mu.Lock()
	updated := append(repo.consignment, consignment)
	repo.consignment = updated
	repo.mu.Unlock()
	return consignment, nil
}
