package main

import (
	"log"
	"net"
	"os"

	"google.golang.org/grpc"

	pb "github.com/RomanGolovinn/KeyManagementSystem/api/proto/kms/v1"

	"github.com/RomanGolovinn/KeyManagementSystem/internal/service"
	"github.com/RomanGolovinn/KeyManagementSystem/internal/storage/postgres"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn, found := os.LookupEnv("DSN")
	if !found {
		log.Fatal("DSN not found")
	}

	psql, err := postgres.NewPostgres(dsn)
	if err != nil {
		log.Fatalf("Failed to create postgres: %v", err)
	}
	defer psql.DB.Close()
	log.Println("Successfully connected to the database!")

	kmsServer := service.NewKMSServer(psql)

	grpcServer := grpc.NewServer()
	pb.RegisterKMSServiceServer(grpcServer, kmsServer)

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on port 50051: %v", err)
	}

	log.Println("Starting gRPC server on :50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve gRPC server: %v", err)
	}
}
