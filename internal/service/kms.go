package service

import (
	"context"

	"github.com/RomanGolovinn/KeyManagementSystem/internal/storage/postgres"

	pb "github.com/RomanGolovinn/KeyManagementSystem/api/proto/kms/v1"
)

type KMSServer struct {
	pb.UnimplementedKMSServiceServer

	db *postgres.Postgres
}

func NewKMSService(db *postgres.Postgres) *KMSServer {
	return &KMSServer{
		db: db,
	}
}

func (s *KMSServer) RegisterClient(ctx context.Context, req *pb.RegisterClientRequest) (*pb.RegisterClientResponse, error) {
	clientModel := postgres.Client{
		Name:   req.GetName(),
		TlsPem: req.GetTlsCert(),
	}

	id, err := s.db.CreateClient(ctx, clientModel)
	if err != nil {
		return nil, err
	}

	return &pb.RegisterClientResponse{
		ClientId: id,
	}, nil
}
