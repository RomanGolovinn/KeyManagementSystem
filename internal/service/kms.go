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

func NewKMSServer(db *postgres.Postgres) *KMSServer {
	return &KMSServer{
		db: db,
	}
}

func (s *KMSServer) RegisterClient(ctx context.Context, req *pb.RegisterClientRequest) (
	*pb.RegisterClientResponse, error) {
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

func (s *KMSServer) RegisterServer(ctx context.Context, req *pb.RegisterServerRequest) (
	*pb.RegisterServerResponse, error) {
	serverModel := postgres.Server{
		Name:    req.GetName(),
		Address: req.GetAddress(),
	}

	id, err := s.db.CreateServer(ctx, serverModel)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterServerResponse{
		ServerId: id,
	}, nil
}

func (s *KMSServer) RegisterPermission(ctx context.Context, req *pb.RegisterPermissionRequest) (
	*pb.RegisterPermissionResponse, error) {
	permissionModel := postgres.Permission{
		ClientId:   req.ClientId,
		ServerId:   req.ServerId,
		Permission: int(req.Permission),
	}
	err := s.db.CreatePermission(ctx, permissionModel)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterPermissionResponse{
		Success: true,
	}, nil
}

func (s *KMSServer) GetSecret(ctx context.Context, req *pb.GetSecretRequest) (*pb.GetSecretResponse, error) {
	// ---

	return &pb.GetSecretResponse{
		Payload: "dummy-secret-for-now",
	}, nil
}
